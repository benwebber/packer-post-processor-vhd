package vhd

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	vboxcommon "github.com/hashicorp/packer/builder/virtualbox/common"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/post-processor/vagrant"
)

// VirtualBoxProvider satisfies the Provider interface.
type VirtualBoxProvider struct {
	name string
}

func NewVirtualBoxProvider() *VirtualBoxProvider {
	return &VirtualBoxProvider{"VirtualBox"}
}

func (p *VirtualBoxProvider) String() string {
	return p.name
}

func (p *VirtualBoxProvider) Name() string {
	return strings.ToLower(p.name)
}

// Execute wraps VBoxManage to run a VirtualBox command.
func (p *VirtualBoxProvider) Execute(ui packer.Ui, command ...string) error {
	driver, err := vboxcommon.NewDriver()
	if err != nil {
		return err
	}
	ui.Message(fmt.Sprintf("Executing: %s", strings.Join(command, " ")))
	if err = driver.VBoxManage(command...); err != nil {
		return err
	}
	return nil
}

// Convert a VirtualBox VMDK artifact to a VHD file.
func (p *VirtualBoxProvider) Convert(ui packer.Ui, artifact packer.Artifact, outputPath string) (err error) {
	var files []string
	// Unpack the VirtualBox artifact if necessary (if in the OVA format).
	for _, path := range artifact.Files() {
		if ext := filepath.Ext(path); ext == ".ova" {
			// Extract OVA files in place.
			ui.Message(fmt.Sprintf("Unpacking OVA: %s", path))
			dir := filepath.Dir(path)
			if err = vagrant.DecompressOva(dir, path); err != nil {
				return err
			}
			// Prepare new slice of files to search.
			glob := filepath.Join(dir, "*")
			files, err = filepath.Glob(glob)
			if err != nil {
				return err
			}
			continue
		} else {
			files = artifact.Files()
		}
	}

	// Find VirtualBox VMDK.
	vmdk, err := findVMDK(files...)
	if err != nil {
		return err
	}
	ui.Message(fmt.Sprintf("Found VirtualBox VMDK: %s", vmdk))

	// Convert VMDK to VHD.
	ui.Message("Cloning VMDK as VHD...")
	command := []string{
		"clonehd",
		"--format", "VHD",
		vmdk,
		outputPath,
	}
	if err = p.Execute(ui, command...); err != nil {
		return fmt.Errorf("Error creating VHD: %s", err)
	}

	return nil
}

// Find the VMDK contained inside the VirtualBox artifact.
func findVMDK(files ...string) (string, error) {
	fileMatches := []string{}
	for _, path := range files {
		if filepath.Ext(path) == ".vmdk" {
			fileMatches = append(fileMatches, path)
		}
	}

	switch len(fileMatches) {
	case 1:
		return fileMatches[0], nil
	case 0:
		return "", errors.New("cannot find VMDK in VirtualBox artifact")
	default:
		return "", errors.New("found multiple VMDKs in VirtualBox artifact")
	}
}

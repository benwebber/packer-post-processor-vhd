package vhd

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	vboxcommon "github.com/mitchellh/packer/builder/virtualbox/common"
	"github.com/mitchellh/packer/packer"
)

// VirtualBoxProvider satisfies the Provider interface.
type VirtualBoxProvider struct{}

func (p *VirtualBoxProvider) String() string {
	return "VirtualBox"
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
func (p *VirtualBoxProvider) Convert(ui packer.Ui, artifact packer.Artifact, outputPath string) error {
	// Find VirtualBox VMDK.
	vmdk, err := findVMDK(artifact.Files()...)
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

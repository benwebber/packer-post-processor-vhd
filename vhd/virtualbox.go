package vhd

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	vboxcommon "github.com/mitchellh/packer/builder/virtualbox/common"
	"github.com/mitchellh/packer/packer"
)

type VirtualBoxProvider struct{}

// Create VHD using VBoxManage.
func (p *VirtualBoxProvider) Convert(ui packer.Ui, artifact packer.Artifact, outputPath string) error {
	// Find VirtualBox VMDK.
	vmdk, err := findVMDK(artifact.Files()...)
	if err != nil {
		return err
	}
	ui.Message(fmt.Sprintf("Found VirtualBox VMDK: %s", vmdk))

	// Convert VMDK to VHD.
	ui.Message("Cloning VMDK as VHD...")
	driver, err := vboxcommon.NewDriver()
	if err != nil {
		return err
	}
	command := []string{
		"clonehd",
		"--format", "VHD",
		vmdk,
		outputPath,
	}
	ui.Message(fmt.Sprintf("Executing: %s", strings.Join(command, " ")))
	if err = driver.VBoxManage(command...); err != nil {
		return fmt.Errorf("Error creating VHD: %s", err)
	}

	return nil
}

// Find the VMDK contained inside the VirtualBox artifact.
func findVMDK(files ...string) (string, error) {
	file_matches := []string{}
	for _, path := range files {
		if filepath.Ext(path) == ".vmdk" {
			file_matches = append(file_matches, path)
		}
	}

	switch len(file_matches) {
	case 1:
		return file_matches[0], nil
	case 0:
		return "", errors.New("cannot find VMDK in VirtualBox artifact")
	default:
		return "", errors.New("found multiple VMDKs in VirtualBox artifact")
	}
}

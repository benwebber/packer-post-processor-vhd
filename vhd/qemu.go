package vhd

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mitchellh/packer/builder/qemu"
	"github.com/mitchellh/packer/packer"
)

type QEMUProvider struct{}

func (p *QEMUProvider) String() string {
	return "QEMU"
}

// PostProcess wraps qemu-img to convert QEMU raw/qcow2 images to VHD files.
func (p *QEMUProvider) Convert(ui packer.Ui, artifact packer.Artifact, outputPath string) error {
	// Find QEMU image.
	img, err := findImage(artifact.Files()...)
	if err != nil {
		return err
	}
	ui.Message(fmt.Sprintf("Found QEMU image: %s", img))

	// Convert image to VHD.
	ui.Message("Converting image to VHD...")
	driver, err := newDriver()
	command := []string{
		"convert",
		"-O", "vpc",
		img,
		outputPath,
	}
	ui.Message(fmt.Sprintf("Executing: %s", strings.Join(command, " ")))
	if err = driver.QemuImg(command...); err != nil {
		return fmt.Errorf("Error creating VHD: %s", err)
	}

	return nil
}

func newDriver() (qemu.Driver, error) {
	qemuImgPath, err := exec.LookPath("qemu-img")
	if err != nil {
		return nil, fmt.Errorf("Failed creating Qemu driver: %s", err)
	}
	driver := &qemu.QemuDriver{
		QemuImgPath: qemuImgPath,
	}
	return driver, nil
}

// Find the image contained inside the QEMU artifact.
func findImage(files ...string) (string, error) {
	file_matches := []string{}
	for _, path := range files {
		if filepath.Ext(path) == ".qcow2" || filepath.Ext(path) == ".raw" {
			file_matches = append(file_matches, path)
		}
	}

	switch len(file_matches) {
	case 1:
		return file_matches[0], nil
	case 0:
		return "", errors.New("cannot find image in QEMU artifact")
	default:
		return "", errors.New("found multiple images in QEMU artifact")
	}
}

// Package vhd implements the packer.PostProcessor interface and adds a
// post-processor that produces a standalone VHD file.
package vhd

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	vboxcommon "github.com/mitchellh/packer/builder/virtualbox/common"
	"github.com/mitchellh/packer/common"
	"github.com/mitchellh/packer/helper/config"
	"github.com/mitchellh/packer/packer"
	"github.com/mitchellh/packer/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`

	OutputPath        string `mapstructure:"output"`
	KeepInputArtifact bool   `mapstructure:"keep_input_artifict"`

	ctx interpolate.Context
}

type PostProcessor struct {
	config Config
}

func (p *PostProcessor) Configure(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		Interpolate: true,
		InterpolateFilter: &interpolate.RenderFilter{
			Exclude: []string{},
		},
	}, raws...)
	if err != nil {
		return err
	}
	return nil
}

// PostProcess wraps VBoxManage to convert a VirtualBox VMDK to VHD file.
func (p *PostProcessor) PostProcess(ui packer.Ui, artifact packer.Artifact) (packer.Artifact, bool, error) {
	// Abort if the input artifact was not produced by the VirtualBox builder.
	switch artifact.BuilderId() {
	case vboxcommon.BuilderId:
		break
	default:
		err := fmt.Errorf("Unknown artifact type: %s\nCan only convert VirtualBox builder artifacts.", artifact.BuilderId())
		return nil, false, err
	}

	ui.Say(fmt.Sprintf("Converting '%s' image to VHD file...", artifact.BuilderId()))

	// Find VirtualBox VMDK.
	vmdk, err := findVMDK(artifact.Files()...)
	if err != nil {
		return nil, false, err
	}
	ui.Message(fmt.Sprintf("Found VirtualBox VMDK: %s", vmdk))

	// Create VHD using VBoxManage.
	artifact = NewArtifact(p.config.OutputPath)
	keep := p.config.KeepInputArtifact

	driver, err := vboxcommon.NewDriver()
	if err != nil {
		return nil, false, err
	}

	ui.Message("Cloning VMDK as VHD...")

	command := []string{
		"clonehd",
		"--format", "VHD",
		vmdk,
		p.config.OutputPath,
	}
	ui.Message(fmt.Sprintf("Executing: %s", strings.Join(command, " ")))
	if err = driver.VBoxManage(command...); err != nil {
		return nil, keep, fmt.Errorf("Error creating VHD: %s", err)
	}

	ui.Say(fmt.Sprintf("Converted VHD: %s", p.config.OutputPath))

	return artifact, keep, nil
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

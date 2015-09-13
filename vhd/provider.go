package vhd

import "github.com/mitchellh/packer/packer"

type Provider interface {
	// Convert converts an artifact into a VHD. The path to the VHD is the
	// third string argument.
	Convert(packer.Ui, packer.Artifact, string) error
}

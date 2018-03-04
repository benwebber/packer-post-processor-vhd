package vhd

import "github.com/hashicorp/packer/packer"

// A Provider wraps logic necessary to convert specific builder artifacts to
// VHD.
type Provider interface {
	// Name should return a simple lowercase identifier for the provider.
	Name() string

	// Execute runs a command using the Provider's Driver.
	Execute(ui packer.Ui, command ...string) error

	// Convert converts a builder artifact into a VHD located at outputPath.
	Convert(ui packer.Ui, artifact packer.Artifact, outputPath string) error

	// String satisfies the Stringer interface and will be used in log
	// messages.
	String() string
}

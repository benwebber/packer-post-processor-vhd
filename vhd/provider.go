package vhd

import "github.com/mitchellh/packer/packer"

// A Provider wraps logic necessary to convert specific builder artifacts to
// VHD.
type Provider interface {
	// Convert converts a builder artifact into a VHD located at outputPath.
	Convert(ui packer.Ui, artifact packer.Artifact, outputPath string) error

	// String satisfies the Stringer interface and will be used in log
	// messages.
	String() string
}

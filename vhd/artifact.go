package vhd

import (
	"fmt"
	"os"
)

// A unique name for this post-processor.
const BuilderId = "benwebber.post-processor.vhd"

// Artifact represents a Virtual Hard Disk (VHD) file.
type Artifact struct {
	// Path is the path to the VHD file on disk.
	Path string
	// Provider represents the upstream source to the
	Provider string
}

// NewArtifact creates a new VHD artifact.
func NewArtifact(provider, path string) *Artifact {
	return &Artifact{
		Path:     path,
		Provider: provider,
	}
}

// BuilderId returns the unique artifact builder ID.
func (a *Artifact) BuilderId() string {
	return BuilderId
}

// Id returns a unique string representing this particular artifact.
func (a *Artifact) Id() string {
	return a.Provider
}

// Files returns a slice of files contained inside the artifact.
func (a *Artifact) Files() []string {
	return []string{a.Path}
}

// String satisfies the Stringer interface.
func (a *Artifact) String() string {
	return fmt.Sprintf("VHD file converted from '%s' image: %s", a.Provider, a.Path)
}

// State represents the state of the VHD artifact and would be used by
// downstream post-processors.
func (*Artifact) State(name string) interface{} {
	return nil
}

// Destroy executes when cleaning up the artifact.
func (a *Artifact) Destroy() error {
	return os.Remove(a.Path)
}

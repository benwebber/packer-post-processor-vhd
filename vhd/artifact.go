package vhd

import (
	"fmt"
	"os"
)

const BuilderId = "benwebber.post-processor.vhd"

type Artifact struct {
	Path string
}

func NewArtifact(path string) *Artifact {
	return &Artifact{
		Path: path,
	}
}

func (a *Artifact) BuilderId() string {
	return BuilderId
}

func (a *Artifact) Id() string {
	return ""
}

func (a *Artifact) Files() []string {
	return []string{a.Path}
}

func (a *Artifact) String() string {
	return fmt.Sprintf("converting: %s", a.Path)
}

func (*Artifact) State(name string) interface{} {
	return nil
}

func (a *Artifact) Destroy() error {
	return os.Remove(a.Path)
}

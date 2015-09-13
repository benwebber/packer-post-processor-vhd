package vhd

import (
	"testing"

	"github.com/mitchellh/packer/packer"
)

func TestPostProcessor_ImplementsPostProcessor(t *testing.T) {
	var _ packer.PostProcessor = new(PostProcessor)
}

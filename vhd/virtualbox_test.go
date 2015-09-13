package vhd

import (
	"testing"
)

func TestVirtualBoxProvider_ImplementsProvider(t *testing.T) {
	var _ Provider = new(VirtualBoxProvider)
}

package vhd

import (
	"testing"
)

func TestQEMUProvider_ImplementsProvider(t *testing.T) {
	var _ Provider = new(QEMUProvider)
}

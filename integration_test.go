// +build integration

package main

import (
	"os"
	"os/exec"
	"testing"
)

var artifacts = []struct {
	in  string
	out string
}{
	{"fixtures/centos-6.7-x86_64.json", "packer_virtualbox-iso_virtualbox.vhd"},
}

func TestIntegration(t *testing.T) {
	if err := os.Chdir("test"); err != nil {
		t.Error(err)
	}
	for _, tt := range artifacts {
		cmd := exec.Command("packer", "build", tt.in)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			t.Error(err)
		}
		if _, err := os.Stat(tt.out); os.IsNotExist(err) {
			t.Error(err)
		}
	}
}

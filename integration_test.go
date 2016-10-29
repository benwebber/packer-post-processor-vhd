// +build integration

package main

import (
	"log"
	"os"
	"os/exec"
	"testing"
)

var artifacts = []struct {
	in  string
	out string
}{
	{"fixtures/virtualbox-ovf.json", "packer_virtualbox-ovf_virtualbox.vhd"},
	{"fixtures/virtualbox-ova.json", "packer_virtualbox-ova_virtualbox.vhd"},
}

// Build the source artifacts.
func init() {
	cmd := exec.Command("packer", "build", "--force", "fixtures/virtualbox-iso.json")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func TestIntegration(t *testing.T) {
	if err := os.Chdir("test"); err != nil {
		t.Error(err)
	}
	for _, tt := range artifacts {
		cmd := exec.Command("packer", "build", "--force", tt.in)
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

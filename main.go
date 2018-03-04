package main

import (
	"github.com/benwebber/packer-post-processor-vhd/vhd"
	"github.com/hashicorp/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(vhd.PostProcessor))
	server.Serve()
}

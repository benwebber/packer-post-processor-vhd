# packer-post-processor-vhd

Packer post-processor plugin to produce Virtual Hard Disk (VHD) files.

VHD files can be used with the following hypervisors:

* Hyper-V
* VMWare
* VirtualBox
* XenServer

## Dependencies

* Packer 0.7+
* VirtualBox (`VBoxManage`)

## Limitations

**packer-post-processor-vhd** currently only supports converting VirtualBox images to VHDs. It can be used as a post-processor for VirtualBox builder artifacts.

## Usage

Add a post-processor declaration to your Packer template:

```json
{
  "post-processors": [
    {
      "type": "vhd",
      "only": ["virtualbox-iso"],
      "output": "builds/centos-6.7-x86_64.vhd"
    }
  ]
}
```

## Installation

1. [Install Packer](https://packer.io/docs/installation.html).

1. Install the [Go toolchain](https://golang.org/doc/install), then install the package:

    ```
    go get github.com/benwebber/packer-post-processor-vhd
    go install github.com/benwebber/packer-post-processor-vhd
    ```

2. Copy the binary to your [Packer plugins directory](https://www.packer.io/docs/extend/plugins.html).

    * Linux and Mac OS X:

        ```
        cp $GOPATH/bin/packer-post-processor-vhd ~/packer.d/plugins
        ```
    * Windows:

        ```
        Copy-Item %GOPATH%\bin\packer-post-processor-vhd %APPDATA%\packer.d\plugins
        ```

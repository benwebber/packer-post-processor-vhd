# packer-post-processor-vhd

Packer post-processor plugin to produce Virtual Hard Disk (VHD) files.

VHD files can be used with the following hypervisors:

* Hyper-V
* VMWare
* VirtualBox
* XenServer

**packer-post-processor-vhd** supports converting [VirtualBox](https://www.packer.io/docs/builders/virtualbox.html) and [QEMU](https://www.packer.io/docs/builders/qemu.html) images to VHDs. It can be used as a post-processor for artifacts from both builders.

## Dependencies

* Packer 0.7+
* VirtualBox (`VBoxManage`) is required to convert VirtualBox artifacts.
* QEMU (`qemu-img`) is required to convert QEMU artifacts.

## Usage

Add a post-processor declaration to your Packer template:

```json
{
  "post-processors": [
    {
      "type": "vhd",
      "only": ["virtualbox-iso"],
    }
  ]
}
```

## Configuration

**packer-post-processor-vhd** supports the following optional configuration items:

* `output` (string)

    The path to the VHD file. This is a [configuration template](https://www.packer.io/docs/templates/configuration-templates.html). The template supports the following variables:

    * `{{ .BuildName }}`

        Replaced by the name of the builder (e.g., `virtualbox-iso` or a custom `name`).

    * `{{ .Provider }}`

        Replaced by the input artifact provider (e.g., `virtualbox`).

    * `{{ .ArtifactId }}`

        Replaced by the ID of the input artifact.

    Defaults to `packer_{{ .BuildName }}_{{ .Provider }}.vhd`.

* `force` (boolean)

    Whether to overwrite a pre-existing VHD file at `output` if it exists. Specifying `--force` on the command line has the same effect. Defaults to `false`.

* `keep_input_artifact` (boolean)

    Whether to keep the input artifact (e.g., VirtualBox image) after processing. Defaults to `false`.

## Installation

### Linux, Mac OS X, and Windows

1. Download the [latest release](https://github.com/benwebber/packer-post-processor-vhd/releases).

2. Rename the plugin `packer-post-processer-vhd`.

3. Copy the binary to your [Packer plugins directory](https://www.packer.io/docs/extend/plugins.html).

    * Linux and Mac OS X:

        ```
        ~/packer.d/plugins
        ```
    * Windows:

        ```
        %APPDATA%\packer.d\plugins
        ```

### Other Platforms


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

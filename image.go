package qemu

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

const (
	ImageFormatRAW   = "raw"
	ImageFormatQCOW2 = "qcow2"
	ImageFormatVDMK  = "vdmk"
	ImageFormatVDI   = "vdi"
	ImageFormatVHDX  = "vhdx"
)

// Image represents a QEMU disk image
type Image struct {
	Path   string // Image location (file)
	Format string // Image format
	Size   uint64 // Image size in bytes

	backingFile string
}

// NewImage constructs a new Image data structure based
// on the specified parameters
func NewImage(path, format string, size uint64) Image {
	var img Image
	img.Path = path
	img.Format = format
	img.Size = size

	return img
}

// LoadImage retreives the information of the specified image
// file into an Image data structure
func LoadImage(path string) (Image, error) {
	type imgInfo struct {
		Format string `json:"format"`
		Size   uint64 `json:"virtual_size"`
	}

	var img Image
	var info imgInfo

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return img, err
	}

	cmd := exec.Command("qemu-img", "info", "--output=json", path)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return img, fmt.Errorf("'qemu-img info' output: %s", oneLine(out))
	}

	err = json.Unmarshal(out, &info)
	if err != nil {
		return img, fmt.Errorf("'qemu-img info' invalid json output: %s", oneLine(out))
	}

	img.Path = path
	img.Format = info.Format
	img.Size = info.Size

	return img, nil
}

// SetBackingFile sets a backing file for the image
// If it is specified, the image will only record the
// differences from the backing file
func (i *Image) SetBackingFile(backingFile string) error {
	if _, err := os.Stat(backingFile); os.IsNotExist(err) {
		return err
	}

	i.backingFile = backingFile
	return nil
}

// Create actually creates the image based on the Image structure
// using the 'qemu-img create' command
func (i Image) Create() error {
	args := []string{"create", "-f", i.Format}

	if len(i.backingFile) > 0 {
		args = append(args, "-o")
		args = append(args, fmt.Sprintf("backing_file=%s", i.backingFile))
	}

	args = append(args, i.Path)
	args = append(args, strconv.FormatUint(i.Size, 10))

	cmd := exec.Command("qemu-img", args...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("'qemu-img create' output: %s", oneLine(out))
	}

	return nil
}

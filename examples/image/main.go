package main

import (
	"fmt"
	"log"

	"github.com/quadrifoglio/go-qemu"
)

const (
	GiB = 1073741824 // 1 GiB = 2^30 bytes
)

func snapshots() {
	img, err := qemu.OpenImage("debian.qcow2")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("base image", img.Path, "format", img.Format, "size", img.Size)

	err = img.CreateSnapshot("backup")
	if err != nil {
		log.Fatal(err)
	}

	snaps, err := img.Snapshots()
	if err != nil {
		log.Fatal(err)
	}

	for _, snapshot := range snaps {
		fmt.Println(snapshot.Name, snapshot.Date)
	}
}

func create() {
	img := qemu.NewImage("vm.qcow2", qemu.ImageFormatQCOW2, 5*GiB)
	img.SetBackingFile("debian.qcow2")

	err := img.Create()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	create()
	snapshots()
}

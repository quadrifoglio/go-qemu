package main

import (
	"github.com/quadrifoglio/go-qemu"

	"fmt"
	"log"
)

func main() {
	img, err := qemu.OpenImage("alpine.qcow2")
	if err != nil {
		log.Fatal(err)
	}

	m := qemu.NewMachine(1, 512)
	m.AddDrive(img)

	pid, err := m.Start("x86_64", true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("QEMU started on PID", pid)
}

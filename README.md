# go-qemu

Golang interface to the QEMU hypervisor

## Installation

```
go get github.com/quadrifoglio/go-qemu
```

You obviously need QEMU to use this tool.

## Usage

### Create an image

```go
img := qemu.NewImage("vm.qcow2", qemu.ImageFormatQCOW2, 5*GiB)
img.SetBackingFile("debian.qcow2")

err := img.Create()
if err != nil {
	log.Fatal(err)
}
```

### Open an existing image

```go
img, err := qemu.OpenImage("debian.qcow2")
if err != nil {
	log.Fatal(err)
}

fmt.Println("image", img.Path, "format", img.Format, "size", img.Size)
```

### Snapshots

```go
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
```

## License

WTFPL (Public Domain)

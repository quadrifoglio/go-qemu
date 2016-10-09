package qemu

import (
	"fmt"
	"os/exec"
	"strconv"
	"syscall"
	"time"
)

// Machine represents a QEMU virtual machine
type Machine struct {
	Cores  int // Number of CPU cores
	Memory int // RAM quantity in megabytes

	drives []Drive
}

// Drive represents a machine hard drive
type Drive struct {
	Path   string // Image file path
	Format string // Image format
}

// NewMachine creates a new virtual machine
// with the specified number of cpu cores and memory
func NewMachine(cores, memory int) Machine {
	var machine Machine
	machine.Cores = cores
	machine.Memory = memory
	machine.drives = make([]Drive, 0)

	return machine
}

// AddDrive attaches a new hard drive to
// the virtual machine
func (m *Machine) AddDrive(image Image) {
	m.drives = append(m.drives, Drive{image.Path, image.Format})
}

// Start stars the machine
// The 'kvm' bool specifies if KVM should be used
// It returns the PID of the QEMU process and an error (if any)
func (m *Machine) Start(arch string, kvm bool) (int, error) {
	qemu := fmt.Sprintf("qemu-system-%s", arch)
	args := []string{"-smp", strconv.Itoa(m.Cores), "-m", strconv.Itoa(m.Memory)}

	if kvm {
		args = append(args, "-enable-kvm")
	}

	for _, drive := range m.drives {
		args = append(args, "-drive")
		args = append(args, fmt.Sprintf("file=%s,format=%s", drive.Path, drive.Format))
	}

	cmd := exec.Command(qemu, args...)
	cmd.SysProcAttr = new(syscall.SysProcAttr)
	cmd.SysProcAttr.Setsid = true

	err := cmd.Start()
	if err != nil {
		return -1, err
	}

	pid := cmd.Process.Pid
	errc := make(chan error)

	go func() {
		err := cmd.Wait()
		if err != nil {
			errc <- fmt.Errorf("'qemu-system-%s': %s", arch, err)
			return
		}
	}()

	time.Sleep(50 * time.Millisecond)

	var vmerr error
	select {
	case vmerr = <-errc:
		if vmerr != nil {
			return -1, vmerr
		}
	default:
		break
	}

	return pid, nil
}

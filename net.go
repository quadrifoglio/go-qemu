package qemu

import (
	"fmt"
)

type NetDev struct {
	Type string // Netdev type (user, tap...)
	ID   string // Netdev ID

	IfName string // TAP: interface name
}

// NewNetworkDevice creates a QEMU network
// device
func NewNetworkDevice(t, id string) (NetDev, error) {
	var netdev NetDev

	if t != "user" && t != "tap" {
		return netdev, fmt.Errorf("Unsupported netdev type")
	}
	if len(id) == 0 {
		return netdev, fmt.Errorf("You must specify a netdev ID")
	}

	netdev.Type = t
	netdev.ID = id

	return netdev, nil
}

// SetHostInterfaceName sets the host interface name
// for the netdev (if supported by netdev type)
func (n *NetDev) SetHostInterfaceName(name string) {
	n.IfName = name
}

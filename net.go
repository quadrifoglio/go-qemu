package qemu

import (
	"fmt"
)

type NetDev struct {
	Type   string // Netdev type (host, tap...)
	ID     string // Netdev ID
	IfName string // Host TAP interface name
}

// NewNetworkDevice creates a QEMU network
// device
func NewNetworkDevice(t, id, ifname string) (NetDev, error) {
	var netdev NetDev

	if t != "user" && t != "tap" {
		return netdev, fmt.Errorf("Unsupported netdev type")
	}
	if len(id) == 0 {
		return netdev, fmt.Errorf("You must specify a netdev ID")
	}

	netdev.Type = t
	netdev.ID = id
	netdev.IfName = ifname

	return netdev, nil
}

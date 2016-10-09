package qemu

import (
	"strings"
)

func oneLine(in []byte) string {
	str := strings.TrimSpace(string(in))
	return strings.Replace(str, "\n", ". ", -1)
}

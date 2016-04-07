package net

import (
	"fmt"
	"strings"
)

const (
	networkDelimiter = "@"
)

// ParseNetwork parse network@(host|ip):port from str.
func ParseNetwork(str string) (network, addr string, err error) {
	idx := strings.Index(str, networkDelimiter)
	if idx == -1 {
		err = fmt.Errorf("addr: \"%s\" error, must be network@(host|ip):port or network@unixsocket", str)
		return
	}
	network = str[:idx]
	addr = str[idx+1:]
	return
}

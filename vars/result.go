package vars

import (
	"net"
)

type Result struct {
	Host     net.IP
	Open     []int
	Closed   []int
	Filtered []int
	MAC      string
	Name     string
}

package types

import (
	"net"
	"github.com/containernetworking/cni/pkg/types"
	"fmt"
)

type HostLocal struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Master string `json:"if0"`
	Mode string `json:"mode"`
	Ipam IPAM `json:"ipam"`
}
type IPAM struct {
	Type string `json:"type"`
	*Range
} 

type Range struct {
	RangeStart net.IP      `json:"rangeStart,omitempty"` // The first ip, inclusive
	RangeEnd   net.IP      `json:"rangeEnd,omitempty"`   // The last ip, inclusive
	Subnet     types.IPNet `json:"subnet"`
	Gateway    net.IP      `json:"gateway,omitempty"`
	Routes [] Route `json:"routes"`
}

type Route struct {
	Dst net.IPNet `json:"dst"`
	GW  net.IP `json:"gw,omitempty"`
}

func (r *Route) String() string {
	return fmt.Sprintf("%+v", *r)
}
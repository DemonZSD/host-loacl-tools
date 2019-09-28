package types

import (
	"fmt"
	"github.com/containernetworking/cni/pkg/types"
	"net"
	"config-writer/utils"
)

type HostLocal struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Master string `json:"if0"`
	Mode   string `json:"mode"`
	Ipam   *IPAM  `json:"ipam"`
}
type IPAM struct {
	Type string `json:"type"`
	*Range
}

type Range struct {
	RangeStart net.IP      `json:"rangeStart,omitempty"` // The first ip, inclusive
	RangeEnd   net.IP      `json:"rangeEnd,omitempty"`   // The last ip, inclusive
	Subnet     *types.IPNet `json:"subnet"`
	Gateway    net.IP      `json:"gateway,omitempty"`
	Routes     []Route     `json:"routes"`
}

type Route struct {
	Dst types.IPNet `json:"dst"`
	GW  net.IP      `json:"gw,omitempty"`
}

func (r *Route) String() string {
	return fmt.Sprintf("%+v", *r)
}

func (rangeIp *Range) SetIpRanges(begain , end net.IP){
	if utils.CompareIp(begain, end) {
		rangeIp.RangeStart = begain
		rangeIp.RangeEnd = end
	}
}
func (rangeIp *Range)SetSubnet(subnet string){
	ipv4Net, err := types.ParseCIDR(subnet)
	if err != nil{
		ipv4Net = nil
	}
	rangeIp.Subnet.IP = ipv4Net.IP
	rangeIp.Subnet.Mask = ipv4Net.Mask
}

func (rangeIp *Range)SetGateway(gateway string){
	rangeIp.Gateway = net.ParseIP(gateway)
}


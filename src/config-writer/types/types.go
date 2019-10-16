package types

import (
	"fmt"
	"github.com/containernetworking/cni/pkg/types"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
)

// 读取 虚拟网卡个数
type VFInfo struct {
	Count  int
	Master string
}
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
	RangeStart net.IP       `json:"rangeStart,omitempty"` // The first ip, inclusive
	RangeEnd   net.IP       `json:"rangeEnd,omitempty"`   // The last ip, inclusive
	Subnet     *types.IPNet `json:"subnet"`
	Gateway    net.IP       `json:"gateway,omitempty"`
	Routes     []Route      `json:"routes"`
}

type Route struct {
	Dst types.IPNet `json:"dst"`
	GW  net.IP      `json:"gw,omitempty"`
}

func (r *Route) String() string {
	return fmt.Sprintf("%+v", *r)
}

func (rangeIp *Range) SetIpRanges(begain, end net.IP) {
	if CompareIp(begain, end) {
		rangeIp.RangeStart = begain
		rangeIp.RangeEnd = end
	}
}
func (rangeIp *Range) SetSubnet(subnet string) {
	ipv4Net, err := types.ParseCIDR(subnet)
	if err != nil {
		ipv4Net = nil
	}
	rangeIp.Subnet.IP = ipv4Net.IP
	rangeIp.Subnet.Mask = ipv4Net.Mask
}

func (rangeIp *Range) SetGateway(subnet string) {
	ipv4Net, err := types.ParseCIDR(subnet)
	if err != nil {
		ipv4Net = nil
	}
	rangeIp.Gateway = ipv4Net.IP
}
func CompareIp(startIp, endIp net.IP) bool {
	startIPArry := strings.Split(startIp.String(), ".")
	endIPArry := strings.Split(endIp.String(), ".")
	startIPNum := make([]int64, 0, 0)
	endIPNum := make([]int64, 0, 0)
	for index, _ := range startIPArry {
		tempStart, _ := strconv.ParseInt(startIPArry[index], 10, 64)
		tempEnd, _ := strconv.ParseInt(endIPArry[index], 10, 64)
		startIPNum = append(startIPNum, tempStart)
		endIPNum = append(endIPNum, tempEnd)
	}

	startIPNumTotal := startIPNum[0]*256*256*256 + startIPNum[1]*256*256 + startIPNum[2]*256 + startIPNum[3]
	endIPNumTotal := endIPNum[0]*256*256*256 + endIPNum[1]*256*256 + endIPNum[2]*256 + endIPNum[3]
	return startIPNumTotal < endIPNumTotal
}

func GetInitIpFromSubset(subset string) (string, error) {
	ipv4Net, err := types.ParseCIDR(subset)
	if err != nil {
		return "", err
	}
	return ipv4Net.IP.String(), nil
}

func (vf *VFInfo) ReadVFNum() (int, error) {
	sriovFile := fmt.Sprintf("/opt/device/sriov_numvfs", vf.Master)
	//sriovFile := fmt.Sprintf("/opt/%s/allocate", vf.Master)
	if _, err := os.Lstat(sriovFile); err != nil {
		return -1, fmt.Errorf("failed to open the sriov_numfs of device %q: %v", vf.Master, err)
	}
	data, err := ioutil.ReadFile(sriovFile)
	if err != nil {
		return -1, fmt.Errorf("failed to read the sriov_numfs of device %q: %v", vf.Master, err)
	}

	if len(data) == 0 {
		return -1, fmt.Errorf("no data in the file %q", sriovFile)
	}
	sriovNumfs := strings.TrimSpace(string(data))
	vfTotal, err := strconv.Atoi(sriovNumfs)
	if err != nil {
		return -1, fmt.Errorf("format num failed %v", err)
	}
	return vfTotal, nil
}

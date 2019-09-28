package utils

import (
	"net"
	"errors"
	"strings"
	"strconv"
)
func CompareIp(startIp, endIp net.IP) bool{
	startIPArry := strings.Split(startIp.String(),".")
	endIPArry := strings.Split(endIp.String(),".")
	startIPNum := make([]int64,0,0)
	endIPNum := make([]int64,0,0)
	for index, _ := range startIPArry  {
		tempStart, _ := strconv.ParseInt(startIPArry[index],10,64)
		tempEnd, _ := strconv.ParseInt(endIPArry[index],10,64)
		startIPNum = append(startIPNum, tempStart)
		endIPNum = append(endIPNum, tempEnd)
	}

	startIPNumTotal := startIPNum[0] * 256 * 256 * 256 + startIPNum[1] * 256 * 256 + startIPNum[2] * 256 + startIPNum[3];
	endIPNumTotal := endIPNum[0] * 256 * 256 * 256 + endIPNum[1] * 256 * 256 + endIPNum[2] * 256 + endIPNum[3];

	return startIPNumTotal < endIPNumTotal
}
func incrementIP(ip net.IP, cidr string) (net.IP, error) {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return ip, err
	}
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] != 0 {
			break
		}
	}
	if !ipNet.Contains(ip) {
		return ip, errors.New("overflowed CIDR while incrementing IP")
	}
	return ip, nil
}


func OffsetIPRange(offset int, endIPV4 net.IP, cidr string)(net.IP, error){
	if offset <= 0 {
		return endIPV4,nil
	}

	for i:=0; i< offset; i++ {
		tempIPV4,err := incrementIP(endIPV4, cidr)
		if err != nil {
			break;
		}
		endIPV4 = tempIPV4
	}
	return endIPV4, nil
}
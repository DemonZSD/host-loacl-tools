package utils

import (
	"net"
	"errors"
)

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
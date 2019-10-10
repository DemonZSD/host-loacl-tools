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


func OffsetIPRange(offset int, initIp net.IP, cidr string)(ip net.IP, err error){
	if offset <= 0 {
		return initIp,nil
	}
	ip = make([]byte, len(initIp))
	copy(ip, initIp)
	for i:=0; i< offset; i++ {
		tempIPV4, err := incrementIP(ip, cidr)
		if err != nil {
			break;
		}
		ip = tempIPV4
	}
	return ip, nil
}
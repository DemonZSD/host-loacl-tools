package utils

import (
	"errors"
	"fmt"
	"net"
)

func incrementIP(ip net.IP, cidr string) (net.IP, error) {
	var initIp net.IP
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return ip, err
	}
	initIp = make([]byte, len(ip))
	copy(initIp, ip)
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] != 0 {
			break
		}
	}
	if !ipNet.Contains(ip) {
		return initIp, errors.New(fmt.Sprintf("overflowed CIDR %s while incrementing IP: %s", cidr, ip.String()))
	}
	return ip, nil
}

func OffsetIPRange(offset int, initIp net.IP, cidr string) (ip net.IP, err error) {
	if offset <= 0 {
		return initIp, nil
	}
	ip = make([]byte, len(initIp))
	copy(ip, initIp)
	for i := 0; i < offset; i++ {
		tempIPV4, err := incrementIP(ip, cidr)
		if err != nil {
			return tempIPV4, err
		}
		ip = tempIPV4
	}
	return ip, nil
}

package common

import (
	"net"
	"strings"
	"tools/iris/contants"
)

func CidrConflict(a, b string) bool {
	for _, cidra := range strings.Split(a, ",") {
		for _, cidrb := range strings.Split(b, ",") {
			if CheckProtocol(cidra) != CheckProtocol(cidrb) {
				continue
			}
			aIp, aIpNet, aErr := net.ParseCIDR(cidra)
			bIp, bIpnet, bErr := net.ParseCIDR(cidrb)
			if aErr != nil || bErr != nil {
				return false
			}
			if aIpNet.Contains(bIp) || bIpnet.Contains(aIp) {
				return true
			}
		}
	}
	return false
}

func CheckProtocol(address string) string {
	ips := strings.Split(address, ",")
	if len(ips) == 2 {
		v4IP := net.ParseIP(strings.Split(ips[0], "/")[0])
		v6IP := net.ParseIP(strings.Split(ips[1], "/")[0])
		if v4IP.To4() != nil && v6IP.To16() != nil {
			return contants.DUAL
		}
		v4IP = net.ParseIP(strings.Split(ips[1], "/")[0])
		v6IP = net.ParseIP(strings.Split(ips[0], "/")[0])
		if v4IP.To4() != nil && v6IP.To16() != nil {
			return contants.DUAL
		}
		return ""
	}
	ip := net.ParseIP(strings.Split(address, "/")[0])
	if ip.To4() != nil {
		return contants.IPV4
	} else if ip.To16() != nil {
		return contants.IPV6
	}
	return ""
}

// ContainsCIDR whether subnet a contains subnet b
func ContainsCIDR(a, b string) bool {
	for _, cidra := range strings.Split(a, ",") {
		for _, cidrb := range strings.Split(b, ",") {
			if CheckProtocol(cidra) != CheckProtocol(cidrb) {
				continue
			}
			_, aIpNet, aErr := net.ParseCIDR(cidra)
			_, bIpnet, bErr := net.ParseCIDR(cidrb)
			if aErr != nil || bErr != nil {
				return false
			}
			aMask, _ := aIpNet.Mask.Size()
			bMask, _ := bIpnet.Mask.Size()
			if aMask <= bMask && aIpNet.Contains(bIpnet.IP) {
				return true
			}
		}
	}
	return false
}

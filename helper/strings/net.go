package strings

import (
	stdStrings "strings"
	"strconv"
	"errors"
	"net"
)

func ParseIpv4(s string) net.IP {
	sPart := stdStrings.Split(s, ".")
	if len(sPart) != net.IPv4len {
		return nil
	}
	ip := make(net.IP, net.IPv4len)
	for i:=0; i<net.IPv4len; i++ {
		if b, err := strconv.Atoi(sPart[i]); err != nil {
			return nil
		} else {
			ip[i] = byte(b)
		}
	}
	return ip
}

func Ip2Uint32(ip string) (result uint32, err error) {
	var (
		ipArr []string
		partN int
	)
	ipArr = stdStrings.Split(ip, ".")
	if len(ipArr) != 4 {
		return 0, errors.New("ip format error")
	}
	for i, ipPart := range ipArr {
		if partN, err = strconv.Atoi(ipPart); partN < 0 || partN > 255 || err != nil {
			return 0, errors.New("ip format error")
		} else {
			result += uint32(partN) << (uint(4-i-1) * 8)
		}
	}
	return
}

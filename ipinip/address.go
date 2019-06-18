package ipinip

import (
	"encoding/binary"
	"net"
)

// IntToIP converts uint32 IP address to net.IP.
func IntToIP(addr uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, addr)
	return ip
}

// IPToInt converts net.IP to uint32.
func IPToInt(addr net.IP) uint32 {
	ip := binary.BigEndian.Uint32(addr.To4())
	return ip
}

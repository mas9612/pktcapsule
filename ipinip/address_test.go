package ipinip

import (
	"net"
	"testing"
)

var intToIPTests = []struct {
	in  uint32
	out net.IP
}{
	{
		3232235521,
		net.IPv4(192, 168, 0, 1),
	},
	{
		2886733835,
		net.IPv4(172, 16, 16, 11),
	},
	{
		167772271,
		net.IPv4(10, 0, 0, 111),
	},
}

func TestIntToIP(t *testing.T) {
	for _, tt := range intToIPTests {
		ip := IntToIP(tt.in)
		if !ip.Equal(tt.out) {
			t.Errorf("IntToIP(%d) = %s, but got %s\n", tt.in, tt.out, ip)
		}
	}
}

var ipToIntTests = []struct {
	in  net.IP
	out uint32
}{
	{
		net.IPv4(192, 168, 0, 1),
		3232235521,
	},
	{
		net.IPv4(172, 16, 16, 11),
		2886733835,
	},
	{
		net.IPv4(10, 0, 0, 111),
		167772271,
	},
}

func TestIPToInt(t *testing.T) {
	for _, tt := range ipToIntTests {
		ip := IPToInt(tt.in)
		if ip != tt.out {
			t.Errorf("IPToInt(%s) = %d, but got %d\n", tt.in, tt.out, ip)
		}
	}
}

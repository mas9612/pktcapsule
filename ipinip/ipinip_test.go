package ipinip

import (
	"bytes"
	"net"
	"testing"
)

var serializeTests = []struct {
	in  IPHeader
	out []byte
}{
	{
		IPHeader{
			Version:        4,
			IHL:            5,
			TotalLength:    0x0127,
			Identification: 0x33b5,
			TimeToLive:     0x2e,
			Protocol:       0x06,
			HeaderChecksum: 0x8c00,
			SrcAddress:     net.IPv4(0x68, 0xf4, 0x2a, 0xc2),
			DstAddress:     net.IPv4(0x0a, 0x61, 0x2e, 0x05),
		},
		[]byte{
			0x45, 0x00, 0x01, 0x27, 0x33, 0xb5, 0x00, 0x00, 0x2e, 0x06, 0x8c, 0x00, 0x68, 0xf4, 0x2a, 0xc2,
			0x0a, 0x61, 0x2e, 0x05,
		},
	},
}

func TestSerialize(t *testing.T) {
	for _, tt := range serializeTests {
		b := tt.in.Serialize()
		if !bytes.Equal(b, tt.out) {
			t.Errorf("IPHeader.Serialize() = % x, but got % x\n", tt.out, b)
		}
	}
}

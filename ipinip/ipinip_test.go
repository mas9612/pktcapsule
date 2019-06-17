package ipinip

import (
	"bytes"
	"net"
	"reflect"
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

var deserializeTests = []struct {
	in  []byte
	out IPHeader
	err error
}{
	{
		[]byte{
			0x45, 0x00, 0x01, 0x27, 0x33, 0xb5, 0x00, 0x00, 0x2e, 0x06, 0x8c, 0x00, 0x68, 0xf4, 0x2a, 0xc2,
			0x0a, 0x61, 0x2e, 0x05,
		},
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
		nil,
	},
}

func TestDeserialize(t *testing.T) {
	for _, tt := range deserializeTests {
		hdr, _ := Deserialize(tt.in)
		if !reflect.DeepEqual(hdr, tt.out) {
			t.Errorf("Deserialize() = %+v, but got %+v\n", tt.out, hdr)
		}
	}
}

var encapsulateTests = []struct {
	in    []byte
	srcIP net.IP
	dstIP net.IP
	out   []byte
}{
	{
		[]byte{
			0x45, 0x00, 0x00, 0x34, 0x00, 0x00, 0x40, 0x00, 0x40, 0x06, 0x6e, 0xa8, 0x0a, 0x61, 0x2e, 0x05,
			0x68, 0xf4, 0x2a, 0xc2, 0xf8, 0xbe, 0x01, 0xbb, 0x5a, 0xb1, 0xd4, 0xa5, 0x5b, 0x7a, 0x35, 0xd7,
			0x80, 0x10, 0x07, 0xec, 0x32, 0x37, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x39, 0x7e, 0xc4, 0x1a,
			0x1d, 0xaa, 0x9a, 0x18,
		},
		net.IPv4(192, 168, 0, 1),
		net.IPv4(192, 168, 0, 2),
		[]byte{
			// Note:
			// Identification and checksum field of outer IP header is 0x00 since they will be computed dynamically and cannot test here.
			// Following test cases are also same as this one.
			0x45, 0x00, 0x00, 0x48, 0x00, 0x00, 0x40, 0x00, 0x3f, 0x04, 0x00, 0x00, 0xc0, 0xa8, 0x00, 0x01,
			0xc0, 0xa8, 0x00, 0x02,
			0x45, 0x00, 0x00, 0x34, 0x00, 0x00, 0x40, 0x00, 0x40, 0x06, 0x6e, 0xa8, 0x0a, 0x61, 0x2e, 0x05,
			0x68, 0xf4, 0x2a, 0xc2, 0xf8, 0xbe, 0x01, 0xbb, 0x5a, 0xb1, 0xd4, 0xa5, 0x5b, 0x7a, 0x35, 0xd7,
			0x80, 0x10, 0x07, 0xec, 0x32, 0x37, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x39, 0x7e, 0xc4, 0x1a,
			0x1d, 0xaa, 0x9a, 0x18,
		},
	},
}

func TestEncapsulate(t *testing.T) {
	for _, tt := range encapsulateTests {
		b, _ := Encapsulate(tt.in, tt.srcIP, tt.dstIP)
		// check before identification field
		if !bytes.Equal(b[:4], tt.out[:4]) {
			t.Errorf("Encapsulate()[:4] = % x, but got % x\n", tt.out[:4], b[:4])
		}
		// check between identification and checksum field
		if !bytes.Equal(b[6:10], tt.out[6:10]) {
			t.Errorf("Encapsulate()[6:10] = % x, but got % x\n", tt.out[6:10], b[6:10])
		}
		// check after checksum field
		if !bytes.Equal(b[12:], tt.out[12:]) {
			t.Errorf("Encapsulate()[12:] = % x, but got % x\n", tt.out[12:], b[12:])
		}
	}
}

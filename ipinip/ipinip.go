package ipinip

import (
	"encoding/binary"
	"net"
)

// IPHeader represents the IPv4 header.
type IPHeader struct {
	// Note: Only lower 4 bits are used.
	Version uint8
	// Note: Only lower 4 bits are used.
	IHL            uint8
	TypeOfService  uint8
	TotalLength    uint16
	Identification uint16
	// Note: Only lower 3 bits are used.
	Flags uint8
	// Note: Higher-3 bits is not used.
	FlagmentOffset uint16
	TimeToLive     uint8
	Protocol       uint8
	HeaderChecksum uint16
	SrcAddress     net.IP
	DstAddress     net.IP
}

// Serialize returns the slice of bytes of IPHeader struct.
// If IHL field is less than 5 (invalid value), empty slice will be returned.
// Currently, IP option is not supported and will be ignored.
func (h IPHeader) Serialize() []byte {
	if h.IHL < 5 {
		return []byte{}
	}

	b := make([]byte, 4*h.IHL)

	b[0] = h.Version<<4 | h.IHL&0x0f
	b[1] = h.TypeOfService
	binary.BigEndian.PutUint16(b[2:], h.TotalLength)
	binary.BigEndian.PutUint16(b[4:], h.Identification)
	b[6] = h.Flags<<5 | uint8(h.FlagmentOffset>>8)&0x1f
	b[7] = uint8(h.FlagmentOffset & 0xff)
	b[8] = h.TimeToLive
	b[9] = h.Protocol
	binary.BigEndian.PutUint16(b[10:], h.HeaderChecksum)
	copy(b[12:], h.SrcAddress.To4())
	copy(b[16:], h.DstAddress.To4())

	return b
}

func Encapsulate(data []byte, srcIP net.IP, dstIP net.IP) []byte {
	return []byte{}
}

func Decapsulate(data []byte, srcIP net.IP, dstIP net.IP) []byte {
	return []byte{}
}

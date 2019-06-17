package ipinip

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/pkg/errors"
)

var (
	ident int32
)

func init() {
	rand.Seed(time.Now().UnixNano())
	ident = rand.Int31n(0x10000)
	fmt.Println(ident)
}

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
//
// HeaderChecksum field will be filled when Serialize is called.
// So you should not assign any value to HeaderChecksum field manually.
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
	copy(b[12:], h.SrcAddress.To4())
	copy(b[16:], h.DstAddress.To4())

	c := checksum(b)
	binary.BigEndian.PutUint16(b[10:], c)

	return b
}

// Deserialize parses given slice of byte and returns the IPHeader struct.
// If some error occured while parsing, empty IPHeader struct and error will be returned.
func Deserialize(data []byte) (IPHeader, error) {
	if len(data) < 20 {
		return IPHeader{}, errors.New("minimum length of IPv4 header is 20 bytes")
	}

	hdr := IPHeader{}
	hdr.Version = data[0] >> 4
	hdr.IHL = data[0] & 0x0f
	hdr.TypeOfService = data[1]
	hdr.TotalLength = binary.BigEndian.Uint16(data[2:])
	hdr.Identification = binary.BigEndian.Uint16(data[4:])
	hdr.Flags = data[6] >> 5
	hdr.FlagmentOffset = uint16(data[6]&0x1f)<<8 | uint16(data[7])
	hdr.TimeToLive = data[8]
	hdr.Protocol = data[9]
	hdr.HeaderChecksum = binary.BigEndian.Uint16(data[10:])
	hdr.SrcAddress = net.IPv4(data[12], data[13], data[14], data[15])
	hdr.DstAddress = net.IPv4(data[16], data[17], data[18], data[19])

	return hdr, nil
}

func identification() uint16 {
	ident++
	if ident > 0xffff {
		ident = 0
	}
	return uint16(ident)
}

// Encapsulate adds newly IP header to the given packet data.
func Encapsulate(data []byte, srcIP net.IP, dstIP net.IP) ([]byte, error) {
	inner, err := Deserialize(data)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Encapsulate failed")
	}

	outer := IPHeader{
		Version:        4,
		IHL:            5,
		TotalLength:    inner.TotalLength + 20,
		Identification: identification(),
		Flags:          FlagDontFragment,
		TimeToLive:     inner.TimeToLive - 1,
		Protocol:       ProtoIP,
		SrcAddress:     srcIP,
		DstAddress:     dstIP,
	}
	outerHdr := outer.Serialize()

	b := make([]byte, outer.TotalLength)
	copy(b, outerHdr)
	copy(b[20:], data)
	return b, nil
}

func Decapsulate(data []byte, srcIP net.IP, dstIP net.IP) []byte {
	return []byte{}
}

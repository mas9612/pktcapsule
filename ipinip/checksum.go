package ipinip

import "encoding/binary"

func checksum(hdr []byte) uint16 {
	var sum uint32
	for i := 0; i < len(hdr); i += 2 {
		if i+2 > len(hdr) { // in the last word, if only one byte remain, add 0x00 as padding
			sum += uint32(binary.BigEndian.Uint16(append(hdr[i:], 0x00)))
		} else {
			sum += uint32(binary.BigEndian.Uint16(hdr[i : i+2]))
		}
		if (sum >> 16) > 0x0 {
			sum += sum >> 16
			sum &= 0xffff // clear carry bit
		}
	}

	return ^uint16(sum) // bit inversion
}

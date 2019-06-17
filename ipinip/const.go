package ipinip

const (
	// FlagDontFragment is the flag which shows the packet must not be fragmented.
	// If this is on, the packet must not be fragmented at the intermediate node.
	FlagDontFragment = 1 << 1
	// FlagMoreFragment is the flag which shows the packet is fragmented and some additional fragmented packet will follow.
	FlagMoreFragment = 1 << 0

	// ProtoIP is the protocol number which is used in IP header.
	// This shows this packet is IPinIP packet and payload is also IP packet.
	ProtoIP = 4
)

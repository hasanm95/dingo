package dns

import "encoding/binary"

// DNS header struct - according to RFC 1035

type DNSHeader struct {
	ID uint16
	QR uint16
	OpCode OpCode
	AA uint16
	TC uint16
	RD uint16
	RA uint16
	Z  uint16
	RCode   ResponseCode
	QDCount uint16
	ANCount uint16
	NSCount uint16
	ARCount uint16
}

type OpCode uint16

const (
	StandardQuery OpCode = 0
)

type ResponseCode uint16

const (
	NoError    ResponseCode = 0
	FormatErr  ResponseCode = 1
)

func (h *DNSHeader) Write() []byte {
	header := make([]byte, 12)

	binary.BigEndian.PutUint16(header, h.ID)

	flags := (h.QR << 15) | (uint16(h.OpCode) << 11) | (h.AA << 10) |
		(h.TC << 9) | (h.RD << 8) | (h.RA << 7) | (h.Z << 4) | uint16(h.RCode)

	binary.BigEndian.PutUint16(header[2:4], flags)
	binary.BigEndian.PutUint16(header[4:6], h.QDCount)
	binary.BigEndian.PutUint16(header[6:8], h.ANCount)
	binary.BigEndian.PutUint16(header[8:10], h.NSCount)
	binary.BigEndian.PutUint16(header[10:12], h.ARCount)

	return header
}
package dns

import (
	"bytes"
	"encoding/binary"
	"strings"
)

type DNSAnswer struct {
	Name  string
	Type  DNSQuestionType
	Class DNSQuestionClass
	TTL   uint32
	Data  string // IP address, যেমন "1.2.3.4"
}

func (a DNSAnswer) Write() []byte {
	var buf bytes.Buffer

	// Name encode
	labels := strings.Split(a.Name, ".")
	for _, label := range labels {
		buf.WriteByte(byte(len(label)))
		buf.WriteString(label)
	}
	buf.WriteByte(0)

	// Type, Class, TTL
	meta := make([]byte, 8)
	binary.BigEndian.PutUint16(meta[0:2], uint16(a.Type))
	binary.BigEndian.PutUint16(meta[2:4], uint16(a.Class))
	binary.BigEndian.PutUint32(meta[4:8], a.TTL)
	buf.Write(meta)

	// RDATA — encode IP address into 4 byte
	rdata := parseIP(a.Data)

	// RDLENGTH
	rdLength := make([]byte, 2)
	binary.BigEndian.PutUint16(rdLength, uint16(len(rdata)))
	buf.Write(rdLength)

	// RDATA
	buf.Write(rdata)

	return buf.Bytes()
}

func parseIP(ip string) []byte {
	parts := strings.Split(ip, ".")
	result := make([]byte, 4)
	for i, part := range parts {
		var num int
		for _, c := range part {
			num = num*10 + int(c-'0')
		}
		result[i] = byte(num)
	}
	return result
}

func WriteAnswers(answers []DNSAnswer) []byte {
	var buf bytes.Buffer
	for _, a := range answers {
		buf.Write(a.Write())
	}
	return buf.Bytes()
}
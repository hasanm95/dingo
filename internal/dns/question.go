package dns

import (
	"bytes"
	"encoding/binary"
	"strings"
)

type DNSQuestionType uint16

const (
	A  DNSQuestionType = 1
	NS DNSQuestionType = 2
)

type DNSQuestionClass uint16

const (
	IN DNSQuestionClass = 1
)

type DNSQuestion struct {
	Name  string
	Type  DNSQuestionType
	Class DNSQuestionClass
}

func (q DNSQuestion) Write() []byte {
	var buf bytes.Buffer

	// encode labels into name: "google.com" → \6google\3com\0
	labels := strings.Split(q.Name, ".")
	for _, label := range labels {
		buf.WriteByte(byte(len(label)))
		buf.WriteString(label)
	}
	buf.WriteByte(0) // null terminator

	// Addition type and class
	typeClass := make([]byte, 4)
	binary.BigEndian.PutUint16(typeClass[0:2], uint16(q.Type))
	binary.BigEndian.PutUint16(typeClass[2:4], uint16(q.Class))
	buf.Write(typeClass)

	return buf.Bytes()
}

func WriteQuestions(questions []DNSQuestion) []byte {
	var buf bytes.Buffer
	for _, q := range questions {
		buf.Write(q.Write())
	}
	return buf.Bytes()
}
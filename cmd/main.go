package main

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/hasanm95/dingo/internal/dns"
)

func main() {
    fmt.Println("Dingo server starting")
    udpAddr, err := net.ResolveUDPAddr("udp", ":2053")

    if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

    conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Dingo DNS server running on port 2053")

    buf := make([]byte, 512)

    for {
        _, source, err := conn.ReadFromUDP(buf)
        if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}

        receivedID := binary.BigEndian.Uint16(buf[0:2])
        
        header := dns.DNSHeader{
			ID:      receivedID,
			QR:      1,
			OpCode:  dns.StandardQuery,
			AA:      0,
			TC:      0,
			RD:      0,
			RA:      0,
			Z:       0,
			RCode:   dns.NoError,
			QDCount: 1,
			ANCount: 1,
			NSCount: 0,
			ARCount: 0,
		}

		// Question
		questions := []dns.DNSQuestion{
			{Name: "codecrafters.io", Type: dns.A, Class: dns.IN},
		}

		// Answer
		answers := []dns.DNSAnswer{
			{Name: "codecrafters.io", Type: dns.A, Class: dns.IN, TTL: 60, Data: "8.8.8.8"},
		}

		// All together
		response := header.Write()
		response = append(response, dns.WriteQuestions(questions)...)
		response = append(response, dns.WriteAnswers(answers)...)

        _, err = conn.WriteToUDP(response, source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
    }
}


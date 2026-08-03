package main

import (
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
        size, source, err := conn.ReadFromUDP(buf)
        if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}

        defaultHeader := dns.DNSHeader{
            ID:      1234,
            QR:      1,
            OpCode:  dns.StandardQuery,
            AA:      0,
            TC:      0,
            RD:      0,
            RA:      0,
            Z:       0,
            RCode:   dns.NoError,
            QDCount: 0,
            ANCount: 0,
            NSCount: 0,
            ARCount: 0,
        }

        receivedData := buf[:size]
		fmt.Printf("Received %d bytes from %s\n", receivedData, source)


		response := defaultHeader.Write()

        _, err = conn.WriteToUDP(response, source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
    }
}


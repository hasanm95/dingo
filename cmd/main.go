package main

import (
	"fmt"
	"net"
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

        receivedData := buf[:size]
		fmt.Printf("Received %d bytes from %s\n", receivedData, source)


		response := []byte("dummy response")

        _, err = conn.WriteToUDP(response, source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
    }
}


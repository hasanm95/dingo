package main

import (
	"fmt"
	"log"
	"net"
)

func main(){
	// 1. Resolving the UDP address
	addr, err := net.ResolveUDPAddr("udp", ":2053")

	if err != nil {
		log.Fatalf("Failed to resolve address: %v", err)
	}

	// 2. Start listening for incoming UDP packets
	conn, err := net.ListenUDP("udp", addr)

	if err != nil {
		log.Fatalf("Failed to start UDP server: %v", err)
	}
	defer conn.Close()

	fmt.Println("Dingo DNS server listening on: 2053")

	// 3. Create buffer to hold incoming packet payloads
	buf := make([]byte, 512)

	// 4. Run infinite loop to receive packages forever
	for {
		// Read incoming packets. Returns bytes read, client info or error
		n, addr, err := conn.ReadFromUDP(buf)

		if err != nil {
			fmt.Printf("Error reading packent: %v\n", err)
			continue
		}

		// Convert bytes to string safely using the length of `n`
		message := string(buf[:n])
		fmt.Printf("Received %d bytes from %s: %s\n", n, addr, message)

		// 5. Echo the message back to the client
		response := []byte("Echo: " + message)
		_, err = conn.WriteToUDP(response, addr)
		if err != nil {
			fmt.Printf("Error sending response to %s: %v\n", addr, err)
		}
	}
}

// func handleQuery(conn *net.UDPConn, buf []byte, remoteAddr *net.UDPAddr) {
// 	fmt.Println()
// }
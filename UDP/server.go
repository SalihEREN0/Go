package main

import (
	"fmt"
	"net"
)

func main() {
	// Create a UDP address to listen on
	addr, err := net.ResolveUDPAddr("udp", ":12345")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// Create a UDP connection
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error creating connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Server is listening on port 12345...")

	buffer := make([]byte, 1024)

	for {
		// Read incoming message
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		message := string(buffer[:n])
		fmt.Printf("Received message: '%s' from %s\n", message, clientAddr)

		// Send response to client
		_, err = conn.WriteToUDP([]byte("Message received"), clientAddr)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
	}
}

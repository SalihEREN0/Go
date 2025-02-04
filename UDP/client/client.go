package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Set up UDP client
	serverAddress := "localhost:8080"
	udpAddr, err := net.ResolveUDPAddr("udp", serverAddress)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Error setting up UDP connection:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Client connected to", serverAddress)

	// Read input from the user and send to server
	for {
		fmt.Print("Enter message (or 'exit' to quit): ")
		var message string
		fmt.Scanln(&message)

		if message == "exit" {
			break
		}

		// Send message to server
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			continue
		}

		// Receive response from server
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading response:", err)
			continue
		}

		// Print the server's response
		fmt.Printf("Server response: %s\n", string(buffer[:n]))
	}
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()
			runClient(clientID)
		}(i)
	}

	wg.Wait()
}

func runClient(clientID int) {
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Client %d: Error resolving server address: %v\n", clientID, err)
		return
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Printf("Client %d: Error connecting to server: %v\n", clientID, err)
		return
	}
	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Client %d: Enter message to send: ", clientID)

		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Client %d: Error reading input: %v\n", clientID, err)
			return
		}

		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Printf("Client %d: Error sending message: %v\n", clientID, err)
			return
		}

		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Client %d: Error reading response: %v\n", clientID, err)
			return
		}

		fmt.Printf("Client %d: Response from server: %s\n", clientID, string(buffer[:n]))
	}
}

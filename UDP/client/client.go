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

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Client %d: Enter message: ", clientID)
		text, _ := reader.ReadString('\n')

		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Printf("Client %d: Error sending message: %v\n", clientID, err)
			return
		}

		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Client %d: Error receiving message: %v\n", clientID, err)
			return
		}

		fmt.Printf("Client %d: Received %d bytes from %s: %s\n", clientID, n, addr, string(buffer[:n]))
	}
}

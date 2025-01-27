package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

func handleConnection(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed")
			return
		}
		fmt.Print("Message from server: " + message)
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ { // Create 5 clients
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()

			conn, err := net.Dial("tcp", "localhost:8080")
			if err != nil {
				fmt.Println("Error connecting:", err)
				return
			}
			defer conn.Close()

			fmt.Fprintf(conn, "Hello from client %d\n", clientID)
			handleConnection(conn, &wg)
		}(i)
	}

	wg.Wait()
}

package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	serverPort := "3000"
	serverAddr := "localhost:" + serverPort

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	defer conn.Close()

	go fmt.Println("Connected:", serverAddr)

	fmt.Println("Enter message:")
	var message string
	fmt.Scanln(&message)

	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal("Message error:", err)
	}

	go fmt.Println("Message sent:", message)
	time.Sleep(1 * time.Second)
}

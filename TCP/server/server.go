package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading from connection: %v\n", err)
		return
	}

	fmt.Println("Received:", string(buffer[:n]))

	_, err = conn.Write([]byte("Message received"))
	if err != nil {
		log.Printf("Error sending response to client: %v\n", err)
		return
	}
}

func main() {
	// ConfigMapten gelen Port bilgisi
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatal("SERVER_PORT environment variable is not set")
	}
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY environment variable is not set")
	}
	fmt.Println("Using secret key:", secretKey)

	serverAddr := ":" + port
	ln, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Error starting server on port %s: %v\n", port, err)
	}
	defer ln.Close()

	fmt.Printf("Server is listening on port %s\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v\n", err)
			continue
		}

		go handleConnection(conn)
	}
}

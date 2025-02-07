package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

func NewServer(port string) {
	port = ":" + port
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("Error starting server on port", port, ":", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server is listening on port", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println("Error reading from connection:", err)
		return
	}

	fmt.Println("Received:", string(buffer[:n]))
}

func main() {
	for i := 0; i < 5; i++ {
		port := strconv.Itoa(3000 + i)
		go NewServer(port)
	}
	select {}
}

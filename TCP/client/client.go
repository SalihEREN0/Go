package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	// ConfigMapten gelen SERVER_ADDR'ı alır
	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		log.Fatal("SERVER_ADDR environment variable is not set")
	}
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Connection error: %v\n", err)
	}
	defer conn.Close()

	fmt.Println("Connected to:", serverAddr)
	go listenForMessages(conn)

	for i := 0; i < 5; i++ { // 5 mesaj örneği
		go sendMessage(conn, fmt.Sprintf("Message %d", i+1))
	}

	//Gönderilnleri bekler
	time.Sleep(2 * time.Second)
}

func sendMessage(conn net.Conn, message string) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		log.Printf("Message error: %v\n", err)
		return
	}
	fmt.Printf("Message sent: %s\n", message)
}

func listenForMessages(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Error reading from server: %v\n", err)
			return
		}
		fmt.Printf("Received from server: %s\n", string(buffer[:n]))
	}
}

package main

import (
	"log"
	"net"
	"sync"
)

type Server struct {
	listenAddr string
	conn       net.PacketConn
	quit       chan struct{}
	clients    map[string]net.Addr
	mu         sync.Mutex
}

func newServer(addr string) *Server {
	return &Server{
		listenAddr: addr,
		quit:       make(chan struct{}),
		clients:    make(map[string]net.Addr),
	}
}

func (s *Server) Start(wg *sync.WaitGroup) error {
	defer wg.Done()

	conn, err := net.ListenPacket("udp", s.listenAddr)
	if err != nil {
		return err
	}
	s.conn = conn
	defer s.conn.Close()

	log.Printf("Server started on %s\n", s.listenAddr)
	go s.listen()

	<-s.quit
	return nil
}

func (s *Server) listen() {
	buffer := make([]byte, 1024)
	for {
		n, addr, err := s.conn.ReadFrom(buffer)
		if err != nil {
			log.Println("Error reading from connection:", err)
			return
		}

		s.addClient(addr)

		log.Printf("Received %d bytes from %s: %s\n", n, addr, string(buffer[:n]))

		s.broadcast(buffer[:n], addr)
	}
}

func (s *Server) addClient(addr net.Addr) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.clients[addr.String()]; !exists {
		log.Printf("New client added: %s\n", addr)
		s.clients[addr.String()] = addr
	}
}

func (s *Server) broadcast(message []byte, sender net.Addr) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, clientAddr := range s.clients {
		if clientAddr.String() != sender.String() {
			_, err := s.conn.WriteTo(message, clientAddr)
			if err != nil {
				log.Printf("Error broadcasting to %s: %v\n", clientAddr, err)
			}
		}
	}
}

func main() {
	var wg sync.WaitGroup

	server := newServer(":1234")

	wg.Add(1)
	go func() {
		if err := server.Start(&wg); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}

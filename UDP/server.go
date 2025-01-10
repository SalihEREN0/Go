package main

import (
	"log"
	"net"
	"sync"
	"time"
)

type Server struct {
	listenAddr string
	conn       net.PacketConn
	quich      chan struct{}
}

func newServer(addr string) *Server {
	return &Server{
		listenAddr: addr,
		quich:      make(chan struct{}),
	}
}

func Start(s *Server, wg *sync.WaitGroup) error {
	defer wg.Done()

	conn, err := net.ListenPacket("udp", s.listenAddr)
	if err != nil {
		return err
	}
	defer conn.Close()
	s.conn = conn
	log.Printf("Server started on %s\n", s.listenAddr)

	go func() {
		time.Sleep(10 * time.Second)
		close(s.quich)
	}()

	<-s.quich
	log.Printf("Server stopped on %s\n", s.listenAddr)
	return nil
}

func main() {
	var wg sync.WaitGroup

	ser := newServer(":1234")
	ser2 := newServer(":1235")
	ser3 := newServer(":1236")

	wg.Add(3)

	go func() {
		if err := Start(ser, &wg); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		if err := Start(ser2, &wg); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		if err := Start(ser3, &wg); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}

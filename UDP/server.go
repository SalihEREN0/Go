package main

import (
	"log"
	"net"
	"time"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	quich      chan struct{}
}

func newServer(addr string) *Server {
	return &Server{
		listenAddr: addr,
		quich:      make(chan struct{}),
	}
}

func Start(s *Server) error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	s.ln = ln

	go func() {
		time.Sleep(5 * time.Second)
		close(s.quich)
	}()

	<-s.quich
	return nil
}

func main() {
	ser := newServer(":1234")
	ser2 := newServer(":1235")
	ser3 := newServer(":1236")
	go log.Fatal(Start(ser))
	time.Sleep(1 * time.Second)
	go log.Fatal(Start(ser2))
	time.Sleep(1 * time.Second)
	go log.Fatal(Start(ser3))
	time.Sleep(1 * time.Second)

	defer ser.ln.Close()
	defer ser2.ln.Close()
	defer ser3.ln.Close()

}

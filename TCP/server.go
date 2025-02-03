package main

import (
	"fmt"
	"net"
)

type Message struct {
	from    string
	payload []byte
}

type Server struct {
	listenAddr string
	ln         net.Listener
	quich      chan struct{}
	msgch      chan Message
}

func newServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quich:      make(chan struct{}),
		msgch:      make(chan Message),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()

	<-s.quich

	close(s.msgch)

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		go s.readLoop(conn)
		fmt.Println("new connection to the server:", conn.RemoteAddr())

	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("reading error:", err)
			return
		}
		msg := buf[:n]
		fmt.Println(string(msg))
	}
}

func main() {
	server := newServer(":3000")
	err := server.Start()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

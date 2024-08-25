package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	quiteChan  chan struct{}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	go s.AcceptLoop()

	<-s.quiteChan
	return nil
}

func (s *Server) AcceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		fmt.Println("new connection:", conn.RemoteAddr())
		go s.ReadLoop(conn)
	}
}

func (s *Server) ReadLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error:", err)
			continue
		}

		msg := buf[:n]
		fmt.Println(string(msg))
	}

}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quiteChan:  make(chan struct{}),
	}
}

func main() {
	server := NewServer(":3000")
	log.Fatal(server.Start())
}

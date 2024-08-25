package main

import (
	"fmt"
	"log"
	"net"
)

type Message struct {
	from    string
	payload []byte
}

type Server struct {
	listenAddr string
	ln         net.Listener
	quiteChan  chan struct{}
	msgChan    chan Message
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
	close(s.msgChan)

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

		s.msgChan <- Message{from: conn.RemoteAddr().String(), payload: buf[:n]}
	}

}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quiteChan:  make(chan struct{}),
		msgChan:    make(chan Message, 1),
	}
}

func main() {
	server := NewServer(":3000")

	go func() {
		for msg := range server.msgChan {
			fmt.Printf("[MESSAGE] (%s):%s\n", string(msg.from), string(msg.payload))
		}
	}()

	log.Fatal(server.Start())
}

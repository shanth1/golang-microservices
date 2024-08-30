package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) HandleWS(ws *websocket.Conn) {
	fmt.Printf("new inconming connection from client: %s\n", ws.RemoteAddr())
	s.conns[ws] = true
	s.ReadLoop(ws)
}

func (s *Server) ReadLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Println("read error:", err)
			continue
		}
		msg := buf[:n]
		s.Broadcast(msg)
	}
}

func (s *Server) Broadcast(b []byte) {
	for w := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := w.Write(b); err != nil {
				fmt.Println("boarcast err:", err)

			}
		}(w)
	}
}

func main() {
	s := NewServer()
	http.Handle("/ws", websocket.Handler(s.HandleWS))
	log.Fatal(http.ListenAndServe(":3000", nil))
}

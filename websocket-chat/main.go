package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
	mu    sync.Mutex
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	defer ws.Close()

	fmt.Printf("New incoming connection from client: %s\n", ws.RemoteAddr())
	s.mu.Lock()
	s.conns[ws] = true
	s.mu.Unlock()

	s.ReadLoop(ws)

	s.removeConn(ws)
}

func (s *Server) removeConn(ws *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.conns, ws)
}

func (s *Server) ReadLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Println("Read error:", err)
			return
		}
		msg := buf[:n]
		s.broadcast(msg)
	}
}

func (s *Server) broadcast(msg []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for ws := range s.conns {
		go func(conn *websocket.Conn) {
			if _, err := conn.Write(msg); err != nil {
				log.Println("Broadcast error:", err)
			}
		}(ws)
	}
}

func (s *Server) startBroadcastTicker() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			message := []byte("Hello from ticker!")
			s.broadcast(message)
		}
	}
}

func main() {
	server := NewServer()
	go server.startBroadcastTicker()

	http.Handle("/ws", websocket.Handler(server.handleWS))
	log.Fatal(http.ListenAndServe(":3000", nil))
}

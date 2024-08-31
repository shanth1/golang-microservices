package main

import (
	"fmt"
	"log"
	"net"
)

type fileServer struct {
}

func (fs *fileServer) start() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go fs.readLoop(conn)
	}
}

func (fs *fileServer) readLoop(conn net.Conn) {
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error on read loop:", err)
		}

		file := buf[:n]
		fmt.Println(file)
		fmt.Printf("received %d bytes\n", n)
	}
}

func main() {
	srv := &fileServer{}
	srv.start()
}

package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net"
	"time"
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

func sendFile(size int) error {
	file := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, file); err != nil {
		return err
	}

	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}

	n, err := conn.Write(file)
	if err != nil {
		return err
	}

	fmt.Printf("written %d bytes\n", n)
	return nil
}

func main() {
	go func() {
		time.Sleep(1 * time.Second)
		if err := sendFile(1000); err != nil {
			fmt.Println("error on sending file:", err)
		}
	}()

	srv := &fileServer{}
	srv.start()

}

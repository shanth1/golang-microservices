package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
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
	buf := new(bytes.Buffer)
	for {
		var size int64
		binary.Read(conn, binary.LittleEndian, &size)
		n, err := io.CopyN(buf, conn, size)
		if err != nil {
			fmt.Println("error on read loop:", err)
		}
		fmt.Println(buf.Bytes())
		fmt.Printf("received %d bytes\n", n)
	}
}

func sendFile(size int64) error {
	file := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, file); err != nil {
		return err
	}

	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}

	binary.Write(conn, binary.LittleEndian, size)

	n, err := io.CopyN(conn, bytes.NewReader(file), size)
	if err != nil {
		return err
	}

	fmt.Printf("written %d bytes\n", n)
	return nil
}

func main() {
	go func() {
		time.Sleep(1 * time.Second)
		if err := sendFile(100000); err != nil {
			fmt.Println("error on sending file:", err)
		}
	}()

	srv := &fileServer{}
	srv.start()

}

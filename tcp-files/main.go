package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

type fileServer struct{}

func (fs *fileServer) start() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen on port 3000: %v", err) // Better error message
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err) // Use log.Printf instead of log.Fatal
			continue
		}
		go fs.readLoop(conn)
	}
}

func (fs *fileServer) readLoop(conn net.Conn) {
	defer conn.Close()

	for {
		var size int64
		err := binary.Read(conn, binary.LittleEndian, &size)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by client")
			} else {
				fmt.Printf("Error reading size from connection: %v\n", err)
			}
			return
		}

		buf := make([]byte, size)
		n, err := io.ReadFull(conn, buf)
		if err != nil {
			fmt.Printf("Error reading data from connection: %v\n", err)
			return
		}

		fmt.Printf("Received %d bytes\n", n)
	}
}

func sendFile(size int64) error {
	file := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, file); err != nil {
		return fmt.Errorf("Failed to generate random data: %v", err)
	}

	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return fmt.Errorf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	err = binary.Write(conn, binary.LittleEndian, size)
	if err != nil {
		return fmt.Errorf("Failed to write size to connection: %v", err)
	}

	n, err := io.CopyN(conn, bytes.NewReader(file), size)
	if err != nil {
		return fmt.Errorf("Failed to write data to connection: %v", err)
	}

	fmt.Printf("Written %d bytes\n", n)
	return nil
}

func main() {
	go func() {
		time.Sleep(1 * time.Second)
		if err := sendFile(100000); err != nil {
			log.Printf("Error sending file: %v", err) // Use log.Printf for error reporting
			os.Exit(1)                                // Exit the process if sending fails
		}
	}()

	srv := &fileServer{}
	srv.start()
}

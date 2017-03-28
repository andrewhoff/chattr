package main

import (
	"fmt"
	"log"
	"net"
)

const (
	addr = "127.0.0.1:7070"
)

func main() {
	fmt.Println("Starting Server...")

	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Listening for connections on %s\n", addr)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error Encountered while accepting a connection: %+v", err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	buf := make([]byte, 128)

	bytesRead, err := conn.Read(buf)
	if err != nil {
		log.Printf("Error Encountered while accepting a connection: %+v", err)
	}

	if bytesRead > 0 {
		fmt.Printf("Received Msg: %s", string(buf))
	}
}

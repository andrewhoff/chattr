package main

import (
	"fmt"
	"log"
	"net"
)

var conns map[net.Conn]string

const addr = "127.0.0.1:7070"

func main() {

	conns = make(map[net.Conn]string)

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
		log.Printf("Registering user: %s", string(buf))
		conns[conn] = string(buf)
	}

	go monitorConn(conn)
}

func monitorConn(conn net.Conn) {
	buf := make([]byte, 128)

	for {
		bytesRead, err := conn.Read(buf)
		if err != nil {
			log.Printf("Error encountered while trying to read from conn: %+v", err)
		}

		if bytesRead > 0 {
			user := conns[conn]
			fmt.Printf("Message received from user: %s. Msg - %s", user, string(buf))
		}
	}
}

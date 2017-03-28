// Written by Andrew Hoff 2017
//
package main

import (
	"fmt"
	"net"
)

const (
	addr = "127.0.0.1:7070"
)

func main() {
	fmt.Println("Running Client..")

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	bytesWritten, err := conn.Write([]byte("hello world!"))
	if err != nil {
		fmt.Printf("Error encountered while attempting to write to conn")
	}

	if bytesWritten > 0 {
		fmt.Println("Message successfully written to connection!")
	}
}

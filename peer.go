// Written by Andrew Hoff 2017
//
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	flag "github.com/ogier/pflag"
)

var username *string = flag.String("username", "Default User", "Enter the username you're signing in as")

const addr = "127.0.0.1:7070"

func main() {
	fmt.Println("Running Peer..")
	flag.Parse()

	if *username == "" {
		log.Fatal("Must have a username to continue")
	}

	fmt.Println("Connecting to server...")
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	// register with server
	err = register(conn)
	if err != nil {
		panic(err)
	}

	bReader := bufio.NewReader(os.Stdin)

	for {
		str, err := bReader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}

		bytesWritten, err := conn.Write([]byte(str))
		if err != nil {
			log.Println("Error encountered while attempting to write to conn")
		}

		if bytesWritten > 0 {
			fmt.Println("Message successfully written to connection!")
		}
	}
}

func register(conn net.Conn) error {
	bytesWritten, err := conn.Write([]byte(*username))
	if err != nil {
		return fmt.Errorf("Error encountered while attempting to write to conn")
	}

	if bytesWritten > 0 {
		fmt.Println("Message successfully written to connection!")
	}

	return nil
}

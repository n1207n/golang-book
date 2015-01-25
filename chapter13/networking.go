package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// Listen on a port
	listener, error := net.Listen("tcp", ":9999")

	if error != nil {
		fmt.Println(error)
		return
	}

	for {
		// Accept a connection
		connection, connectionError := listener.Accept()

		if connectionError != nil {
			fmt.Println(connectionError)
			continue
		}

		// Handle the connection
		go handlerServerConnection(connection)
	}
}

func handlerServerConnection(connection net.Conn) {
	// Receive the message
	var msg string

	decodeError := gob.NewDecoder(connection).Decode(&msg)

	if decodeError != nil {
		fmt.Println(decodeError)
	} else {
		fmt.Println("Received", msg)
	}

	connection.Close()
}

func client() {
	// Connect to the server
	connection, connectionError := net.Dial("tcp", "127.0.0.1:9999")

	if connectionError != nil {
		fmt.Println(connectionError)
		return
	}

	// Send the message
	msg := "Hello World"
	fmt.Println("Sending", msg)

	encodeError := gob.NewEncoder(connection).Encode(msg)

	if encodeError != nil {
		fmt.Println(encodeError)
	}

	connection.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}

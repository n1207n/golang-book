package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Server struct{}

func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

func server() {
	rpc.Register(new(Server))

	listener, listenerError := net.Listen("tcp", ":9999")

	if listenerError != nil {
		fmt.Println(listenerError)
		return
	}

	for {
		connection, connectionError := listener.Accept()

		if connectionError != nil {
			continue
		}

		go rpc.ServeConn(connection)
	}
}

func client() {
	connection, connectionError := rpc.Dial("tcp", "127.0.0.1:9999")

	if connectionError != nil {
		fmt.Println(connectionError)
		return
	}

	var result int64

	connectionCallError := connection.Call("Server.Negate", int64(999), &result)

	if connectionError != nil {
		fmt.Println(connectionCallError)
	} else {
		fmt.Println("Server.Negate(999) = ", result)
	}
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}

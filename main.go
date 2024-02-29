package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"net-cat/functions"
)

var mutex sync.Mutex

func main() {
	port := "8989"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	server, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}

	defer server.Close()
	fmt.Printf("Listening on the port %s\n", port)
	go functions.Broadcast(&mutex)
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			conn.Close()
			continue
		}
		go functions.Handle(conn, &mutex)
	}
}

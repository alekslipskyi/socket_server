package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))

	if err != nil {
		fmt.Printf("error is %s", err)
	}

	defer listener.Close()

	fmt.Printf("Server runned on port %s", os.Getenv("PORT"))

	for {
		conn, err := listener.Accept()

		writer := bufio.NewWriter(conn)
		if err != nil {
			log.Printf("error accepting connection %v", err)
			continue
		}
		log.Printf("accepted connection from %v", conn.RemoteAddr())


	}
}

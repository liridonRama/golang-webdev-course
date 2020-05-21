package main

import (
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		con, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handle(con)
	}
}

func handle(con net.Conn) {
	io.WriteString(con, "I see you connected")
	con.Close()
}
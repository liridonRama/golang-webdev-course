package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		con, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handle(con)
	}

}

func handle(con net.Conn) {
	defer con.Close()
	con.SetDeadline(time.Now().Add(time.Second * 10))
	reader := bufio.NewScanner(con)

	for reader.Scan() {
		fmt.Println(reader.Text())
	}

	fmt.Println("Code got here.")
	io.WriteString(con, "I see you connected.")
}

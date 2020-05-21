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
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handle(conn)
	}

}

func handle(con net.Conn) {
	defer con.Close()
	con.SetDeadline(time.Now().Add(time.Second * 10))
	reader := bufio.NewScanner(con)

	for reader.Scan() {
		ln := reader.Text()
		fmt.Println(ln)

		if ln == "" {
			break
		}
	}

	body := "<h1>LOVE IS NICE</h1>"

	io.WriteString(con, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(con, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(con, "Content-Type: text/html\r\n")
	io.WriteString(con, "\r\n")
	io.WriteString(con, body)
}

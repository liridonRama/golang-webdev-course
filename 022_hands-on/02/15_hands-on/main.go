package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
	reader, counter := bufio.NewScanner(con), 0
	var method, uri string

	for reader.Scan() {
		ln := reader.Text()
		fmt.Println(ln)

		if counter == 0 {
			xs := strings.Split(ln, " ")
			method, uri = xs[0], xs[1]
		}

		if ln == "" {
			break
		}
		counter++
	}

	body := "<h1>HOLY COW THIS IS LOW LEVEL</h1>"

	io.WriteString(con, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(con, "METHOD: %v\r\n", method)
	fmt.Fprintf(con, "URI: %v\r\n", uri)
	fmt.Fprintf(con, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(con, "Content-Type: text/html\r\n")
	fmt.Fprint(con, "Method: text/html\r\n")
	io.WriteString(con, "\r\n")
	io.WriteString(con, body)
}

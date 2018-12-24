package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	var i int
	var rMethod string
	var rURI string
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			pt := strings.Fields(ln)
			rMethod = pt[0]
			rURI = pt[1]
			fmt.Println("Method:" + rMethod)
			fmt.Println("URI: " + rURI)
		}
		if ln == "" {
			break
		}
		fmt.Println(ln)
		i++
	}
	body := "I saw you"
	body += "\nMethod: " + rMethod
	body += "\nURI: " + rURI
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

}

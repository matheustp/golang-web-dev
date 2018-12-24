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
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Code Gangsta</title>
	</head>
	<body>
		<h1>"HOLY COW THIS IS LOW LEVEL"</h1>
	</body>
	</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

}

package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
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
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			break
		}
		fmt.Println(ln)
	}

	io.WriteString(conn, "I saw you")

}
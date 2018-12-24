package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	response(conn, request(conn))
}

func request(conn net.Conn) string {
	i := 0
	scanner := bufio.NewScanner(conn)
	var url string
	for scanner.Scan() {
		ln := scanner.Text()

		fmt.Println(ln)
		if i == 0 {
			url = strings.Fields(ln)[1]
			fmt.Println("*** URL", url)
		}
		if ln == "" {
			break
		}
		i++
	}
	return url
}
func response(conn net.Conn, url string) {

	//body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>%v</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(url))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, url)
}

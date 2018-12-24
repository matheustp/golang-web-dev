package main

import (
	"bufio"
	"fmt"
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
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			if ln == "" {
				break
			}
			fmt.Println(ln)

		}
		defer conn.Close()
		fmt.Println("I see you")
		conn.Write([]byte("I see you"))
		//or 		io.WriteString(conn, "I see you connected.")

		conn.Close()
	}
}

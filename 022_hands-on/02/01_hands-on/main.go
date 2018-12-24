package main

import "net"

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
		conn.Write([]byte("I see you"))
		//or 		io.WriteString(conn, "I see you connected.")

		conn.Close()
	}
}

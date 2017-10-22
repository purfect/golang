package main

import (
	"fmt"
	"io"
	"time"
	"net"
)

func main() {
	newserver, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer newserver.Close()

	for {
		conn, err := newserver.Accept()
		if err != nil {
			panic(err)
		}
		io.WriteString(conn, fmt.Sprint("Hello Server\n", time.Now(), "\n"))
		conn.Close()
	}
}

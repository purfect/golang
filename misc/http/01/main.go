package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		io.WriteString(conn, "\n[ TCP -  Server ]\n")
		fmt.Fprintln(conn, "Das ist ein Test!")
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
}

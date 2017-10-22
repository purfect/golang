package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const ADDRESS = ":9000"

func main() {
	ln, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		fmt.Println("Starten des Servers fehlgeschlagen:", err.Error())
		os.Exit(1)
	}
	// channel für die eingehenden Verbindungen, Nachrichten und "tote" Verbindungen
	aconns := make(map[net.Conn]int)
	conns := make(chan net.Conn)
	dconns := make(chan net.Conn)
	msgs := make(chan string)
	i := 0

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println("Fehler: ", err.Error())
			}
			conns <- conn
			fmt.Println("Verbunden: ",conn.RemoteAddr(), " (Client-Nr.:", i+1,")")
		}
	}()
	for {
		select {
		// lesen der eingehenden Verbindungen
		case conn := <-conns:
			aconns[conn] = i
			i++
			//sobald die Verbindung aktiv ist wird davon gelsen
			go func(conn net.Conn, i int) {
				rd := bufio.NewReader(conn)
				for {
					m, err := rd.ReadString('\n')
					if err != nil {
						break
					}
					msgs <- fmt.Sprintf("Client %v: %v", i, m)
				}
				// fertig mit lesen
				dconns <- conn
			}(conn, i)
		case msg := <-msgs:
			// ausgeben der Nachrichten an alle
			for conn := range aconns {
				conn.Write([]byte(msg))
			}
		case dconn := <-dconns:
			fmt.Printf("Client %v hat die Verbindung getrennt\n", aconns[dconn]+1)
			delete(aconns, dconn)
		}
	}
}

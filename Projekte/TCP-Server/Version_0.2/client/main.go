package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main () {
	// Die Verbindung zum Server
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("Es konnte keine Verbindung zum Server hergestellt werden:", err)
	}
	for{
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Eingabe: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Es ist ein Fehler bei der Eingabe aufgetreten:", err)
			os.Exit(1)
		}
		// Daten übertragen
		fmt.Fprintf(conn, text + "\n")
		// empfangen der Antwort
		message, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Antwort: " + message)
	}
}

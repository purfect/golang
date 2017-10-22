package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"strings"
)

func main ()  {
	fmt.Println("Starte Server... (Beenden mit Strg + c)")
	// wir lauschen auf allen interfaces auf Port 9000
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Der konnte nicht gestartet werden:", err.Error())
		os.Exit(1)
	}
	// akzeptieren der Verbindungen zum Port
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Es wurde eine Verbindung hergestellt ")
	// eine unendliche Schleife (Abbruch mit Strg + c)
	for {
		// warten auf ende der Übertragung mit newline (\n)
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Es trat ein Fehler bei der Verbindung auf:", err.Error())
		}
		// Ausgabe der Übertragenen Nachricht
		fmt.Println("nachricht erhalten:", string(message))
		// OK-Meldung des Servers über den erhalt der Nachricht
		newMessage := "OK:"+strings.ToUpper(message)
		conn.Write([]byte(newMessage+"\n"))
	}
}

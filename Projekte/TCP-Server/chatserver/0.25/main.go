package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"io"
	"time"
)

type User struct {
	Name   string
	Output chan Message
}

type Message struct {
	Username string
	Text     string
}
type ChatServer struct {
	Users map[string]User
	Join  chan User
	Leave chan User
	Input chan Message
}

func (cs *ChatServer) Run() {
	for {
		select {
		case user := <-cs.Join:
			cs.Users[user.Name] = user
			go func() {
				cs.Input <- Message{
					Username: "System",
					Text:     fmt.Sprintf("%s ist verbunden", user.Name),
				}
			}()
		case user := <-cs.Leave:
			delete(cs.Users, user.Name)
			go func() {
				cs.Input <- Message{
					Username: "System",
					Text:     fmt.Sprintf("%s hat die Verbindung getrennt", user.Name),
				}
			}()
		case msg := <-cs.Input:
			for _, user := range cs.Users {
				select {
					case user.Output <- msg:
				default:

				}
			}
		}
	}
}

func handleConn(chatServer *ChatServer, conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "Username:")

	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	user := User{
		Name:   scanner.Text(),
		Output: make(chan Message,10),
	}
	chatServer.Join <- user
	defer func() {
		chatServer.Leave <- user
	}()

	// Read from conn
	go func() {
		for scanner.Scan(){
			ln := scanner.Text()
			chatServer.Input <- Message{user.Name, ln}
		}
	}()

	// write to conn
	for msg := range user.Output {
		if msg.Username != user.Name {
			_, err := io.WriteString(conn, "[" + time.Now().Format(time.RFC822) + "] " + msg.Username + ": " + msg.Text+"\n")
			if err != nil {
				break
			}
		}
	}
}

func main() {
	server, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	defer server.Close()

	chatServer := &ChatServer{
		Users: make(map[string]User),
		Join:  make(chan User),
		Leave: make(chan User),
		Input: make(chan Message),
	}
	go chatServer.Run()
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Error:", err.Error())
			os.Exit(1)
		}
		fmt.Println("[" + time.Now().Format(time.RFC822) + "] es wurde eine neue Verbindung hergestellt")
		go handleConn(chatServer, conn)
	}

}

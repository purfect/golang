package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://httpbin.org/basic-auth/user/passwßrd")
	if err != nil {
		log.Fatalln("Keine Antwort erhalten")
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Konnte Inhalt nicht lesen")
	}
	fmt.Println("Inhalt:", string(content))
	fmt.Println("Statuscode:", resp.StatusCode)
}

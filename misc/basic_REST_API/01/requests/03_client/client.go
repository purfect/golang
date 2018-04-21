package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)

	if err != nil {
		log.Fatalln("Es konnte keine Anfrage erstellt werden ")
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Keine Antwort erhalten")
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Inhalt konnte nicht gelesen werden")
	}
	fmt.Println(string(content))
}

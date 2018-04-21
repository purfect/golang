package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln("Keine Antwort")
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Inhalt konnte nicht gelesen werden")
	}
	fmt.Println(string(content))
}

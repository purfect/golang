package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Post("https://httpbin.org/post", "text/plain",
		strings.NewReader("das ist der angefragte Inhalt"))
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

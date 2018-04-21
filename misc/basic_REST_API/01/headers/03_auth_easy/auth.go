package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "https://httpbin.org/basic-auth/user/passw0rd", nil)
	if err != nil {
		log.Fatalln("es konnte keine Anfrage erstellt werden")
	}
	req.SetBasicAuth("user", "passw0rd")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("konnte Antwort nicht lesen")
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Konnte Inhalt nicht lesen")
	}
	fmt.Println("Inhalt:", string(content))
	fmt.Println("Statuscode:", resp.StatusCode)
}

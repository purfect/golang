package main

import (
	"bytes"
	"encoding/base64"
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
	buffer := &bytes.Buffer{}
	enc := base64.NewEncoder(base64.URLEncoding, buffer)
	enc.Write([]byte("user:passw0rd"))
	enc.Close()
	encodeCreds, err := buffer.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		log.Fatalln("Failed to read encoded buffer")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodeCreds))
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

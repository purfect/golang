package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Es fehlt ein Dateiname")
		os.Exit(1)
	}
	file := arguments[1]
	// Get bytes from file
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println()
	fmt.Println("Datei: ", file)
	fmt.Println()
	// Hash the file and output results
	fmt.Printf("Md5: %x\n\n", md5.Sum(data))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(data))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(data))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(data))
}

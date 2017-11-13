package main

import (
	"fmt"
	"math/big"
	"crypto/rand"
	"os"
)

func main() {
	rand, err := rand.Int(rand.Reader, big.NewInt(1984))
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Printf("Random Number: %d\n", rand)
}

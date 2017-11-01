package main

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r1.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	fmt.Println(randSeq(10))
}

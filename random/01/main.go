package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 10; i++ {
		fmt.Print("Random-Number: ")
		fmt.Println(r1.Intn(100))
	}
}

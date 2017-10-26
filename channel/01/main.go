package main

import (
	"fmt"
	"math/rand"
	"time"
)
const count = 10

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)
	go send(even, odd, quit)
	receive(even, odd, quit)
}
func send(even, odd chan<- int, quit chan<- bool) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < count; i++ {
		checkint := r1.Intn(100)
		if checkint%2 == 0 {
			even <- checkint
		} else {
			odd <- checkint
		}
	}
	quit <- true
}
func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("gerade: ",v)
		case v := <-odd:
			fmt.Println("ungerade: ", v)
		case <-quit:
			return
		}
	}
}

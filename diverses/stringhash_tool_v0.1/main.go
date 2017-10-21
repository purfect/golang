package main

import (
	"flag"
	"fmt"
	"github.com/jzelinskie/whirlpool"
	"os"
)

func main(){
	stringoption := flag.String("s", "", "an string, -s=\"test test2\"")
	flag.Parse()
	if len(*stringoption) < 1  {
		fmt.Println("es wurde kein String-Argument angegeben!")
		os.Exit(1)
	}
	w := whirlpool.New()
	hash := []byte(*stringoption)
	w.Write(hash)
	fmt.Println("Argument: ", *stringoption)
	fmt.Printf("Hash: %x\n",  w.Sum(nil))
}

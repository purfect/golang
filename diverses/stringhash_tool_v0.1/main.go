package main

import (
	"flag"
	"fmt"
)

func main(){
	stringoption := flag.String("s", "", "an string, -s=\"test test2\"")
	flag.Parse()
	fmt.Println(*stringoption)
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}

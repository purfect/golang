package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Es fehlt ein Argument!")
		os.Exit(1)
	}
	file := arguments[1]
	info, err := os.Stat(file)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	mode := info.Mode()
	filesize := info.Size()
	fmt.Println(file, ": ", mode)
	fmt.Println("Size: ", filesize/1024, "kb")
}

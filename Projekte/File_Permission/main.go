/*
* Programm zeigt Permissions an
* 
* $ go run main.go /bin/watch
* /bin/watch :  -rwxr-xr-x
*
* Coder: Rasputin
* Date: 21.10.2017
*
*/
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
	fmt.Println(file, ": ", mode, "\n")
}

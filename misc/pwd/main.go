/* 
* Implementierung des pwd-Befehls in Go
*
* Coder: Rasputin
* Date: 21.10.2017
*
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	arguments := os.Args
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	} else {
		fmt.Println(pwd)
	}
	if len(arguments) == 1 {
		return
	}
	if arguments[1] != "-l" {
		return
	}
	fileinfo, err := os.Lstat(pwd)
	if fileinfo.Mode()&os.ModeSymlink != 0 {
		realpath, err := filepath.EvalSymlinks(pwd)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		} else {
			fmt.Println(realpath)
		}
	}
}

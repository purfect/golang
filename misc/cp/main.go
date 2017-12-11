package main

import (
	"os"
	"fmt"
	"io"
	"path/filepath"
	"strconv"
)

var BUFFERSIZE int64

func Copy(src, dst string, BUFFERSIZE int64) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		fmt.Printl("[!] Error:", err)
		os.Exit(1)
	}
	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("[!] Error: %s is not a regular file!\n", err)
		os.Exit(1)
	}
	source, err := os.Open(src)
	if err != nil {
		fmt.Println("[!] Error: ", err)
		os.Exit(1)
	}
	defer source.Close()

	_, err := os.Stat(dst)
	if err != nil {
		fmt.Println("[!] Error: ", err)
		os.Exit(1)
	}
	destination, err := os.Create(dst)
	if err != nil {
		fmt.Println("[!] Error: ", err)
		os.Exit(1)
	}
	defer destination.Close()

	if err != nil {
		panic(err)
	}
	buf := make([]byte, BUFFERSIZE)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("[!] Error: ", err)
			os.Exit(1)
		}
		if n == 0 {
			break
		}
		_, err := destination.Write(buf[:n])
		if err != nil) {
			return err
		}
	}
	return err
}

func main()  {
	if len(os.Args) != 4 {
		fmt.Printf("[*] usage: %s source destination BUFFERSIZE\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	source := os.Args[1]
	destination := os.Args[2]
	BUFFERSIZE, _ = strconv.ParseInt(os.Args[3], 10,64)
	fmt.Printf("[*] Copying %s to %s\n", source, destination)
	err := Copy(source, destination, BUFFERSIZE)
	if err != nil {
		fmt.Printf("[!] File copying failed: %s", err)
	}
}

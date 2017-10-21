package main

import (
	"flag"
	"fmt"
	"github.com/jzelinskie/whirlpool"
	"os"
	"encoding/hex"
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
	// check
	fmt.Println("Check")
	hashstring := "b913d5bbb8e461c2c5961cbe0edcdadfd29f068225ceb37da6defcf89849368f8c6c2eb6a4c4ac75775d032a0ecfdfe8550573062b653fe92fc7b8fb3b7be8d6"
	checkstring := hex.EncodeToString(w.Sum(nil))
	fmt.Println(checkstring)
	if checkstring == hashstring {
		fmt.Println("Hash passt")
	} else {
		fmt.Println("etwas passt nicht")
		os.Exit(1)
	}
}

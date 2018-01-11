package main
import (
	"os"
	"fmt"
	"github.com/beevik/ntp"
) 

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(time)
}

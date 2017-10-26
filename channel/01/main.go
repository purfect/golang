/*
* Hier wird eine bestimmte Anzahl an Zufalszahlen generiert
* und diese werden geprüft ob sie durch 2 teilbar sind oder nicht
* Anschliessend werden diese auf entsprechende Channels verteilt
*
* Coder: Rasputin
* Date: 26.10.2017
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)
const count = 10

func main() {
	// Die Channels
	even := make(chan int)
	odd := make(chan int)
	//quit wird zum beenden der Schleife beim empfangen der Werte aus den Channeln benötigt
	quit := make(chan bool)
	// Funktion zum Senden der zufallsnummern in den Channel
	go send(even, odd, quit)
	// Funktion zum Empfangen der zufallsnummern aus den Channeln
	receive(even, odd, quit)
}
func send(even, odd chan<- int, quit chan<- bool) {
	// erzeugen der Zufallsnummern
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// Zufälliger Integer zwischen 0 - 1000 wird bei jedem Aufruf in der Schleife erzeugt
	// und geprüft ob dieser durch 2 teilbar ist
	// gerade Werte gehen in den even-channel
	// ungerade Werte gehen in den odd-channel
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
// empfangen und ausgeben der Daten
func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Printf("gerade:\t\t%d\n",v)
		case v := <-odd:
			fmt.Printf("ungerade:\t%d\n", v)
		case <-quit:
			fmt.Println("Ende")
			return
		}
	}
}

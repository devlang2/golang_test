package main

import (
	"log"
	"time"
)

func main() {
	//

	go func() {
		i := 0
		for {
			log.Print(i)
			i++
			time.Sleep(1 * time.Second)
		}
	}()

	timer1 := time.NewTimer(time.Second * 3)
	<-timer1.C
	log.Println("Timer 1 expired")
}

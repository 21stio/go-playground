package main

import (
	"log"
	"time"
)

func main() {
	chan1 := make(chan string, 100)
	chan1<-"chan1_1"

	chan2 := make(chan string, 100)
	chan2<-"chan2_1"

	var usedChannel *chan string = &chan1

	go func() {
		for {
			channel := *usedChannel
			msg := <-channel
			log.Print(msg)
			channel <- msg
		}
	}()

	time.Sleep(2 * time.Second)

	usedChannel = &chan2

	time.Sleep(3 * time.Second)
}

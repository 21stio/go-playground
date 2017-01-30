package main

import (
	"log"
	"time"
)

func main() {
	chan1 := make(chan string, 100)
	chan1<-"chan1_1"
	chan1<-"chan1_2"
	chan1<-"chan1_3"
	chan1<-"chan1_4"
	chan1<-"chan1_5"

	chan2 := make(chan string, 100)
	chan2<-"chan2_1"
	chan2<-"chan2_2"
	chan2<-"chan2_3"
	chan2<-"chan2_4"
	chan2<-"chan2_5"

	usedChannel := chan1

	go func() {
		for {
			log.Print(<-usedChannel)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(2 * time.Second)

	usedChannel = chan2

	time.Sleep(3 * time.Second)
}

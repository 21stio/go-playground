package main

import (
	"log"
	"time"
	"sync"
)

func closeAndSwitchChannel()  {
	chan1 := make(chan string, 100)
	chan1<-"1"

	chan2 := make(chan string, 100)
	chan2<-"2"

	var usedChannel *chan string = &chan1

	lock := sync.Mutex{}

	go func() {
		for {
			time.Sleep(100 * time.Millisecond)

			lock.Lock()
			channel := *usedChannel

			msg := <-channel
			log.Print(msg)

			channel <- msg
			lock.Unlock()
		}
	}()

	time.Sleep(2 * time.Second)

	lock.Lock()

	time.Sleep(1 * time.Second)

	usedChannel = &chan2

	lock.Unlock()


	time.Sleep(3 * time.Second)
}

func switchChannel()  {
	chan1 := make(chan string, 100)
	chan1<-"1"

	chan2 := make(chan string, 100)
	chan2<-"2"

	var usedChannel *chan string = &chan1

	go func() {
		for {
			timeout := make(chan bool, 1)
			go func() {
				time.Sleep(100 * time.Millisecond)
				timeout <- true
			}()

			select {
			case msg := <-*usedChannel:
				log.Print(msg)
			case <-timeout:
				continue
			}
		}
	}()

	time.Sleep(2 * time.Second)

	usedChannel = &chan2

	time.Sleep(3 * time.Second)
}

func channelBlocking()  {
	chan1 := make(chan bool)

	chan1 <- true
}

func main() {
	channelBlocking()
}
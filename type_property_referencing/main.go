package main

import (
	"k8s.io/contrib/compare/Godeps/_workspace/src/github.com/emicklei/go-restful/log"
	"time"
)

type Whatever struct{
	Name string
}

func (whatever Whatever) SetNameByCopy(name string) {
	log.Printf("Whatever.SetNameByCopy() memory address: %p", whatever)

	whatever.Name = name
}

func (whatever Whatever) GetNameByCopy() (string) {
	log.Printf("Whatever.GetNameByCopy() memory address: %p", whatever)

	return whatever.Name
}

func (whatever *Whatever) SetNameByReference(name string) {
	log.Printf("Whatever.SetNameByReference() memory address: %p", whatever)

	whatever.Name = name
}

func (whatever *Whatever) GetNameByReference() (string) {
	log.Printf("Whatever.GetNameByReference() memory address: %p", whatever)

	return whatever.Name
}

func (whatever *Whatever) LogNameEverySecond() () {
	log.Printf("Whatever.LogNameEverySecond() memory address: %p", whatever)

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			<-ticker.C
			log.Printf("Whatever.LogNameEverySecond() GetNameByCopy(): %v", whatever.GetNameByCopy())
			log.Printf("Whatever.LogNameEverySecond() GetNameByReference(): %v", whatever.GetNameByReference())
		}
	}()
}

func evaluateNameByCopy()  {
	whatever := Whatever{}

	whatever.SetNameByCopy("bzz")

	name := whatever.GetNameByCopy()

	log.Printf("evaluateNameByCopy name: %v", name)
}

func evaluateMemoryAddressWhenNotWritingAnything()  {
	whatever := Whatever{}

	whatever.GetNameByReference()
	whatever.GetNameByReference()
	whatever.GetNameByReference()
}

func evaluateNameByReference()  {
	whatever := Whatever{}

	whatever.SetNameByReference("bzz")

	name := whatever.GetNameByReference()

	log.Printf("evaluateNameByReference name: %v", name)
}

func evaluateAccessFromWithinAGoRoutine()  {
	whatever := Whatever{}

	whatever.LogNameEverySecond()

	time.Sleep(2 * time.Second)

	whatever.SetNameByReference("bzz")

	stop := make(chan bool)
	<-stop
}

func main() {
	evaluateNameByReference()
}
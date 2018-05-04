package main

import (
	"log"
)

func abc(a ...interface{}) {
	b, ok := a[0].(string)
	if ! ok {
		log.Fatal()
	}

	c, ok := a[1].(string)
	if ! ok {
		log.Fatal()
	}

	println(b)
	println(c)
}

func main()  {
	abc("a")
}
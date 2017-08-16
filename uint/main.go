package main

import "fmt"

func main()  {
	run()
}

func run() {
	a := uint64(1000)
	b := uint64(500)

	fmt.Printf("%v", b - a)
}
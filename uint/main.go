package main

import "fmt"

func main()  {
	testNegativeIntToUint()
}

func run() {
	a := uint64(1000)
	b := uint64(500)

	fmt.Printf("%v", b - a)
}

func testNegativeIntToUint() {
	a := int64(-1000)
	b := uint64(a)

	fmt.Printf("%v", b)
}
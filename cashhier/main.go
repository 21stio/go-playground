package main

import (
	"fmt"
	"os"
)

func fact(n int64) int64 {
	if n == 0 {
		return 1
	}
	i := n * fact(n - 1)
	return int64(i)
}

func main() {
	fmt.Printf("Halo Krazimir %s \n", os.Args[1])
        fmt.Printf("N Factorial is: %v \n", fact(7))
}
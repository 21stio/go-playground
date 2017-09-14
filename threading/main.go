package main

import (
	"runtime"
	"log"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (error) {
	print(runtime.NumCPU())

	return nil
}

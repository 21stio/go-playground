package main

import "github.com/davecgh/go-spew/spew"

func main() {
	data := []string{"a", "b", "c"}
	//ixs := []int{2,1}

	spew.Dump(data[0:1:2])
}

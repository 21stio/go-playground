package main

import "github.com/davecgh/go-spew/spew"

func main() {
	m := map[string]*string{}

	hi := "hi"
	m["a"] = &hi

	a := m["a"]

	ho := "ho"
	*m["a"] = ho

	spew.Dump(a)
}

package main

import (
	"regexp"
	"github.com/davecgh/go-spew/spew"
)

func main() {

	r := regexp.MustCompile(`^[A-Z][a-zA-Z]*$`)

	a := "AffeAffe1"
	//b := "affeAffe"

	s := r.FindString(a)

	spew.Dump(s)

	//spew.Dump(ms)
}

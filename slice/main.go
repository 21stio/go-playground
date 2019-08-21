package main

import (
	"github.com/davecgh/go-spew/spew"
	"strings"
)

func Title(ss []string) (ss0 []string) {
	for _, s := range ss {
		ss0 = append(ss0, strings.Title(s))
	}

	return
}

func Concat(ss []string) (s string) {
	for _, v := range ss {
		s += v
	}

	return
}

func main() {
	a := []string{"la","le","li","lu","lo",}

	a = append(a[:1], Title(a[1:])...)

	spew.Dump(a)
}

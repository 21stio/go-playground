package main

import (
	"reflect"
	"github.com/davecgh/go-spew/spew"
)

type A struct {
	AString string
}

type B struct {
	BString string
}

func main() {
	var any interface{}
	any = A{}

	spew.Dump(reflect.TypeOf(any))
}



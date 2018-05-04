package main

import (
	"github.com/davecgh/go-spew/spew"
	"reflect"
)

type AnyStruct struct {
	A string
}

func (a AnyStruct) String(b string) string {
	return b
}

func main() {
	a := (AnyStruct).String(AnyStruct{}, "hi")

	b := (AnyStruct)(AnyStruct{A:"yo"}).A

	reflect.TypeOf((AnyStruct)(AnyStruct{A:"yo"}).A)

	spew.Dump(reflect.TypeOf((AnyStruct{}).A))

	spew.Dump(a, b)
}
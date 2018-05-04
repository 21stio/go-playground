package main

import (
	"github.com/davecgh/go-spew/spew"
	"encoding/binary"
	"unsafe"
	"fmt"
)

type Msg struct {
	a  *string
	b  *string
	c  *string
	d  *string
	aa *string
	ba *string
	ca *string
	da *string
	Strings *Strings
}

type Http struct {
}

type Strings struct {
	a string
	b string
	c string
	d string
}

type Int struct {
	a string
	b string
	c string
	d string
}

func main() {
	txt := "halloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasd"
	m := Msg{}
	fmt.Printf("sizeof(Coord3d) = %d\n", unsafe.Sizeof(txt))

	//m.Strings = Strings{}
	fmt.Printf("sizeof(Coord3d) = %d\n", unsafe.Sizeof(m))

	//r := reflect.ValueOf("halloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasdhalloasdasdasdasdsadasd")
	s := binary.Size(txt)

	spew.Dump(s)
}

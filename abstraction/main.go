package main

import "github.com/davecgh/go-spew/spew"

type Field struct {
	Name string
}

func (f Field) GetName() (string) {
	return f.Name
}

type Whatever struct {
	Field
	Name string
}


type Lala []string

func (l *Lala) Append(ss []string) {
	l0 := append(*l, ss...)
	l = &l0
}

type LoLo map[string]string

func (l LoLo) Append(k, v string) {
	l[k] = v
}

func main() {
	l := LoLo{}

	l.Append("a", "b")

	spew.Dump(l)
}

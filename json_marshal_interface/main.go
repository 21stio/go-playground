package main

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"log"
)

type Wrapper struct {
	Object interface{} `json:"object"`
}

type classTrait struct {
	Class string `json:"@class"`
}

type Cat struct {
	classTrait
	Name string `json:"name"`
}

func NewCat() (c Cat) {
	c.Class = "Cat"

	return
}

type Car struct {
	classTrait
	Speed string `json:"speed"`
}

func NewCar() (c Car) {
	c.Class = "Car"

	return
}

func main() {
	err := func() (err error) {
		cat := NewCat()
		cat.Name = "Bob"

		w := Wrapper{}
		w.Object = cat

		b, err := json.Marshal(w)
		if err != nil {
			return
		}

		spew.Dump(b)

		nc := NewCat()
		nw := Wrapper{Object: &nc}

		err = json.Unmarshal(b, &nw)
		if err != nil {
			return
		}

		spew.Dump(nc)

		return
	}()
	if err != nil {
		log.Fatal(err)
	}
}

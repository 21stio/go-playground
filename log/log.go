package main

import (
	"fmt"
	"log"
	"errors"
)

type Yo struct {

}

func (y Yo) String() string {
	return "yoyasdsdoy"
}

func main()  {
	errors.New()
	a := Yo{}
	fmt.Print(a)

}

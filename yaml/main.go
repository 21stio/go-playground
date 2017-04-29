package main

import (
	"github.com/ghodss/yaml"
	"fmt"
)

type Person struct {
	Name string `json:"name"`  // Affects YAML field names too.
	Age int `json:"age"`
}

func main() {
	p := Person{"John", 30}

	y, err := yaml.Marshal(p)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	print(string(y))
	print("aa")
}

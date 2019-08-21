package main

import (
	"github.com/davecgh/go-spew/spew"
	"reflect"
)

func removeEmpty(v reflect.Value) (reflect.Value) {
	t := v.Type()

	for n := 0; n < t.NumField(); n++ {
		v0 := v.Field(n)
		if v0.Kind() != reflect.Ptr {
			continue
		}

		spew.Dump(v0.Elem().Kind())
		if v0.Elem().Kind() != reflect.Struct {
			continue
		}

		v0.Elem().Set(removeEmpty(v0.Elem()))

		if isEmpty(v0.Elem()) {
			v0.Elem().Set(reflect.Zero(v0.Elem().Type()))
		}
	}

	return v
}

func isEmpty(v reflect.Value) (bool) {
	t := v.Type()

	for n := 0; n < t.NumField(); n++ {
		if !v.Field(n).IsNil() {
			return false
		}
	}

	return true
}

type Person struct {
	Name *string `json:"name,omitempty"`
	Mom  *Person`json:"mom,omitempty"`
	Dad  *Person`json:"dad,omitempty"`
}


func stringPtr(s string) (*string) {
	return &s
}

func main() {
	p:= Person{
		Name: stringPtr("Justus"),
		Dad: &Person{
			Name: stringPtr("Peter"),
		},
		Mom: &Person{
			Name: stringPtr("Matilda"),
			Dad: &Person{},
		},
	}

	p0 := removeEmpty(reflect.ValueOf(p)).Interface().(Person)

	spew.Dump(p0)


}

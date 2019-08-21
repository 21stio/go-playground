package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"reflect"
)

func structsToColumns(ss interface{}) (columns map[string][]string) {
	structColumns := []map[string]string{}

	v := reflect.ValueOf(ss)

	for i := 0; i < v.Len(); i++ {
		structColumns = append(structColumns, structToColumns("", v.Index(i).Interface()))
	}

	columnNames := map[string]bool{}
	for _, c := range structColumns {
		for k, _ := range c {
			columnNames[k] = true
		}
	}

	columns = map[string][]string{}
	for _, c := range structColumns {
		for k, _ := range columnNames {
			v := c[k]
			columns[k] = append(columns[k], v)
		}
	}

	return
}

func structToColumns(prefix string, s interface{}) (columns map[string]string) {
	if prefix != "" {
		prefix += "."
	}

	spew.Dump(s)

	columns = map[string]string{}

	v := reflect.ValueOf(s)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		f := v.Field(i)
		fName := prefix + t.Field(i).Name

		if f.IsNil() {
			continue
		}

		switch f.Elem().Kind() {
		case reflect.Struct:
			columns0 := structToColumns(fName, f.Elem().Interface())
			for k, v := range columns0 {
				columns[k] = v
			}
		case reflect.Int32:
			columns[fName] = fmt.Sprintf("%v", f.Elem().Int())
		case reflect.Float64:
			columns[fName] = fmt.Sprintf("%v", f.Elem().Float())
		case reflect.Bool:
			columns[fName] = fmt.Sprintf("%v", f.Elem().Bool())
		case reflect.String:
			columns[fName] = f.Elem().String()
		}
	}

	return
}

type Person struct {
	Name   *string
	Father *Person
	Mother *Person
}

func stringPtr(s string) (*string) {
	return &s
}

func main() {
	ps := []Person{
		{
			Name: stringPtr("Justus"),
			Mother: &Person{
				Name: stringPtr("Matilda"),
				Mother: &Person{
					Name: stringPtr("GrandMatilda"),
				},
			},
			Father: &Person{
				Name: stringPtr("Titus"),
			},
		},
		{
			Name: stringPtr("Peter"),
			Mother: &Person{
				Name: stringPtr("Jenny"),
			},
			Father: &Person{
				Name: stringPtr("Baba"),
			},
		},
		{
			Name: stringPtr("Bob"),
		},
	}

	c := structsToColumns(ps)

	spew.Dump(c)
}

package main

import (
	"github.com/davecgh/go-spew/spew"
	"reflect"
	"github.com/mitchellh/mapstructure"
)

type Whatever struct {
	StringField *string
	IntField    *int32
	FloatField  *float64
	BoolField   *bool
	ObjectField *Whatever
	StringList  []string
	IntList     []int32
	FloatList   []float64
	BoolList    []bool
	ObjectList  []Whatever
}

type Generic struct {
	Type        string
	String      map[string]string
	Int32       map[string]int32
	Float64     map[string]float64
	Bool        map[string]bool
	StringList  map[string][]string
	Int32List   map[string][]int32
	Float64List map[string][]float64
	BoolList    map[string][]bool
	Generic     map[string]Generic
	GenericList map[string]GenericList
}

type Enum struct {
	Type  string
	Value int32
}

type EnumList struct {
	Type string
	List []Enum
}

type GenericList struct {
	Type string
	List []Generic
}

func main() {
	stringField := "a"
	intField := int32(1)
	floatField := 1.0
	boolField := true

	whatever := Whatever{
		StringField: &stringField,
		IntField:    &intField,
		FloatField:  &floatField,
		BoolField:   &boolField,
		StringList:  []string{"a"},
		IntList:     []int32{1},
		FloatList:   []float64{1},
		BoolList:    []bool{true},
		ObjectField: &Whatever{
			StringField: &stringField,
			IntField:    &intField,
			FloatField:  &floatField,
			BoolField:   &boolField,
			StringList:  []string{"a"},
			IntList:     []int32{1},
			FloatList:   []float64{1},
			BoolList:    []bool{true},
		},
		ObjectList: []Whatever{
			Whatever{
				StringField: &stringField,
				IntField:    &intField,
				FloatField:  &floatField,
				BoolField:   &boolField,
				StringList:  []string{"a"},
				IntList:     []int32{1},
				FloatList:   []float64{1},
				BoolList:    []bool{true},
			},
			Whatever{
				StringField: &stringField,
				IntField:    &intField,
				FloatField:  &floatField,
				BoolField:   &boolField,
				StringList:  []string{"a"},
				IntList:     []int32{1},
				FloatList:   []float64{1},
				BoolList:    []bool{true},
			},
		},
	}

	g := toGeneric(whatever)

	spew.Dump(g)

	n := fromGeneric(g)

	var w Whatever
	mapstructure.Decode(n, &w)

	spew.Dump(w)
}

func toGeneric(i interface{}) (g Generic) {
	v := reflect.ValueOf(i)

	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("not struct")
	}

	t := v.Type()

	g.Type = t.Name()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.PkgPath != "" {
			continue
		}

		v0 := v.FieldByIndex([]int{i})

		if v0.IsNil() {
			continue
		}

		switch v0.Kind() {
		case reflect.Ptr:
			switch v0.Elem().Kind() {
			case reflect.String:
				if g.String == nil {
					g.String = map[string]string{}
				}

				g.String[field.Name] = v0.Elem().Interface().(string)

				break
			case reflect.Int32:
				if g.Int32 == nil {
					g.Int32 = map[string]int32{}
				}

				g.Int32[field.Name] = v0.Elem().Interface().(int32)

				break
			case reflect.Float64:
				if g.Float64 == nil {
					g.Float64 = map[string]float64{}
				}

				g.Float64[field.Name] = v0.Elem().Interface().(float64)

				break
			case reflect.Bool:
				if g.Bool == nil {
					g.Bool = map[string]bool{}
				}

				g.Bool[field.Name] = v0.Elem().Interface().(bool)

				break
			case reflect.Struct:
				if g.Generic == nil {
					g.Generic = map[string]Generic{}
				}

				g.Generic[field.Name] = toGeneric(v0.Elem().Interface())

				break
			}
		case reflect.Slice:
			t0 := reflect.TypeOf(v0.Interface())
			switch t0.Elem().Kind() {
			case reflect.String:
				if g.StringList == nil {
					g.StringList = map[string][]string{}
				}

				g.StringList[field.Name] = v0.Interface().([]string)

				break
			case reflect.Int32:
				if g.Int32List == nil {
					g.Int32List = map[string][]int32{}
				}

				g.Int32List[field.Name] = v0.Interface().([]int32)

				break
			case reflect.Float64:
				if g.Float64List == nil {
					g.Float64List = map[string][]float64{}
				}

				g.Float64List[field.Name] = v0.Interface().([]float64)

				break
			case reflect.Bool:
				if g.BoolList == nil {
					g.BoolList = map[string][]bool{}
				}

				g.BoolList[field.Name] = v0.Interface().([]bool)

				break
			case reflect.Struct:
				if g.GenericList == nil {
					g.GenericList = map[string]GenericList{}
				}

				l := GenericList{}
				l.Type = t0.Name()

				for i := 0; i < v0.Len(); i++ {
					l.List = append(l.List, toGeneric(v0.Index(i).Interface()))
				}

				g.GenericList[field.Name] = l

				break
			}
		}
	}

	return
}

func fromGeneric(g Generic) (m map[string]interface{}) {
	m = map[string]interface{}{}

	for k, v := range g.String {
		m[k] = v
	}

	for k, v := range g.StringList {
		m[k] = v
	}

	for k, v := range g.Int32 {
		m[k] = v
	}

	for k, v := range g.Int32List {
		m[k] = v
	}

	for k, v := range g.Float64 {
		m[k] = v
	}

	for k, v := range g.Float64List {
		m[k] = v
	}

	for k, v := range g.Bool {
		m[k] = v
	}

	for k, v := range g.BoolList {
		m[k] = v
	}

	for k, v := range g.Generic {
		m[k] = fromGeneric(v)
	}

	for k, v := range g.GenericList {
		l := []map[string]interface{}{}
		for _, v0 := range v.List {
			l = append(l, fromGeneric(v0))
		}

		m[k] = l
	}

	return
}

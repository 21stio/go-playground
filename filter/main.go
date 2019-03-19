package main

import (
	"github.com/21stio/go-playground/filter/types"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

type StringFilter struct {
	Contains *string
	Set *bool
}

type FilterScope string

const (
	NONE  FilterScope = "none"
	SOME              = "some"
	EVERY             = "every"
)

type PersonFilter struct {
	Name    *StringFilter
	Mother  *PersonFilter
	Friends *PersonsFilter
}

type PersonsFilter struct {
	Scope FilterScope
	Name  *StringFilter
}

func FilterPersons(filter PersonFilter, persons types.Persons) (persons0 []types.Person) {
	persons = types.Persons(persons)
	persons.SetIndices()

	indices := filterPersons(filter, persons)

	return persons.Select(indices)
}

func filterPersons(filter PersonFilter, persons types.Persons) ([]int) {
	if filter.Name != nil {
		indices := filterStrings(*filter.Name, persons.GetNames())

		persons = persons.Select(indices)
	}

	return persons.GetIndices()
}

func filterStrings(filter StringFilter, strings0 types.Strings) ([]int) {
	if filter.Contains != nil {
		filtered := types.Strings{}

		for _, v := range strings0 {
			if v.Value != nil && strings.Contains(*v.Value, *filter.Contains) {
				filtered = append(filtered, v)
			}
		}

		strings0 = filtered
	}

	if filter.Set != nil {
		filtered := types.Strings{}

		if *filter.Set == true {
			for _, v := range strings0 {
				if v.Value != nil {
					filtered = append(filtered, v)
				}
			}
		} else {
			for _, v := range strings0 {
				if v.Value == nil {
					filtered = append(filtered, v)
				}
			}
		}

		strings0 = filtered
	}

	return strings0.GetIndices()
}

func StringPtr(s string) (*string) {
	return &s
}

func main() {
	persons := types.Persons{
		{
			Name: &types.String{Value:StringPtr("justus")},
			Mother: &types.Person{
				Name: &types.String{Value:StringPtr("justina")},
			},
			Friends: types.Persons{

			},
		},
		{
			Name: &types.String{Value:StringPtr("peter")},
			Mother: &types.Person{
				Name: &types.String{Value:StringPtr("petrara")},
			},
		},
		{
			Name: &types.String{Value:StringPtr("bob")},
			Mother: &types.Person{
				Name: &types.String{Value:StringPtr("bobba")},
			},
		},
	}

	filter := PersonFilter{
		Name: &StringFilter{
			Contains: StringPtr("bob"),
		},
	}

	persons = FilterPersons(filter, persons)

	spew.Dump(persons)
}

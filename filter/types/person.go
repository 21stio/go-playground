package types

type Person struct {
	_Index  int
	Name    *String
	Mother 	*Person
	Friends []Person
}

type Persons []Person

func (persons Persons) Select(indices []int) (persons0 Persons) {
	for _, i := range indices {
		persons0 = append(persons0, persons[i])
	}

	return
}

func (persons Persons) SetIndices() () {
	for i, _ := range persons {
		persons[i]._Index = i
	}

	return
}

func (persons Persons) GetIndices() (indices []int) {
	for i, _ := range persons {
		indices = append(indices, persons[i]._Index)
	}

	return
}

func (persons Persons) GetNames() (names []String) {
	for i, _ := range persons {
		name := *persons[i].Name
		name._Index = i

		names = append(names, name)
	}

	return
}

func (persons Persons) GetMothers() (mothers Persons) {
	for i, _ := range persons {
		mother := *persons[i].Mother
		mother._Index = i

		mothers = append(mothers, mother)
	}

	return
}

func (persons Persons) GetFriendss() (friends []Persons) {
	for i, _ := range persons {
		friends = append(friends, persons[i].Friends)
	}

	return
}
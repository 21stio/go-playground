package types

type String struct {
	_Index  int
	Value *string
}

type Strings []String

func (strings0 Strings) GetIndices() (indices []int) {
	for i, _ := range strings0 {
		indices = append(indices, strings0[i]._Index)
	}

	return
}

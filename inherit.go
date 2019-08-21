package main


type CreateTrait struct {
	v string
}

func (t *CreateTrait) Create(v string) {
	t.v = v
}

type DeleteTrait struct {
	v string
}

func (t *DeleteTrait) Create(v string) {
	t.v = v
}

func main() {
	
}

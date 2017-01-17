package main

type IBaseRepository interface {
	GetModels() ([]BaseModel)
}

type BaseModel struct {
	Name string
}

type PetRepository struct {
}

func (repository PetRepository) GetModels() ([]BaseModel)  {
	return []PetModel{
		PetModel{
			Name: "Carl",
		},
		PetModel{
			Name: "Fritz",
		},
		PetModel{
			Name: "Peter",
		},
	}
}

type PetModel struct {
	BaseModel
	Name string
}

func (model PetModel) GetName() (string)  {
	return "chris"
}

func PrintNames(repository IBaseRepository) () {
	for _, model := range repository.GetModels() {
		print(model.GetName())
	}
}

func PrintName(model BaseModel) () {
	print(model.GetName())
}

func main() {
	abc := PetModel{
		Name: "Karl",
	}

	PrintName(abc)
}

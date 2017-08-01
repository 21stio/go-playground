package main

func main() {
	defer1()
}

func defer0()  {
	color := "green"
	defer println(*&color)
	color = "blue"
}

func defer1()  {
	color := "green"
	defer func(){println(*&color)}()
	color = "blue"
}

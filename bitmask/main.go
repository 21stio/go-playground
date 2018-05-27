package main

const (
	A = 1 << iota
	B
	C
	D
	E
)


func main() {
	//var abc int = 0

	abc := E|A

	if (abc&A) != 0 {
		println("A")
	}

	if (abc&B) != 0 {
		println("B")
	}

	if (abc&C) != 0 {
		println("C")
	}

	if (abc&D) != 0 {
		println("D")
	}

	if (abc&E) != 0 {
		println("E")
	}
}

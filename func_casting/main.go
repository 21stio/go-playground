package main

type StringFunc func(string) (string)


func main() {
	var abc interface{} = func(s string) (string) {
		return s
	}

	switch abc.(type) {
	case StringFunc:
		println("yo")
	case func(string) (string):
		println("yo1")
	}
}

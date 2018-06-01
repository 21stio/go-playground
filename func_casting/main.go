package main

type StringFunc func(string) (string)

func getStringFunc() StringFunc {
	return func(s string) (string) {
		return s
	}
}

func main() {
	var abc interface{} = getStringFunc()

	switch abc.(type) {
	case StringFunc:
		println("yo")
	case func(string) (string):
		println("yo1")
	}
}

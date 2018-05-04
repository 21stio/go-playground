package main

import "time"

func main() {
	m := map[string]map[string]map[string]bool{}
	println(m["a"]["b"]["c"])

	do()
}

func do(sth interface{})  {
	p := ""

	switch sth.(type) {
	case string:
		p = "string"
		break
	case []byte:
		p = "byte"
		break
	case time.Duration:
		p = "time"
		break
	}

	println(p)
}



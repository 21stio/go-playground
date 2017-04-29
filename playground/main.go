package main

import (
	"strings"
)

func indent(spaces int, v string) string {
	pad := strings.Repeat(" ", spaces)
	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}

func main() {
	print(indent(2, "yo:\n  hi"))
	print("\nabc")
}

package main

import "github.com/davecgh/go-spew/spew"

func main() {
	err := pickingNumbers()
	if err != nil {
	    panic(err)
	}
}

func pickingNumbers() (err error) {
	ns := []int{66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66}

	cs := map[int]int{}

	for _, n := range ns {
		cs[n]++
	}

	addedCs := map[int]int{}

	greatest := 0
	for k, c := range cs {
		c0, ok := cs[k+1]
		if !ok {
		    continue
		}

		added := c + c0

		if added > greatest {
			greatest = added
		}

		addedCs[k] = added
	}

	spew.Dump(addedCs)

	return
}

package main

import "github.com/davecgh/go-spew/spew"

func main() {
	spew.Dump(find("apple", "ale"))
	spew.Dump(find("apple", "aep"))
}

func find(haystack, needle string) (bool) {
	nIndex := 0
	if len(needle) == 0 {
		return true
	}

	for i:=0; i < len(haystack); i++ {
		if haystack[i] == needle[nIndex] {
			nIndex += 1
			if nIndex == len(needle) {
				return true
			}
		}
	}

	return false
}

func intersection () {
	a := []string{"a", "b", "c"}
	b := []string{"b", "c", "d"}

	m := map[string]bool{}
	intersection := []string{}
	for _, s := range a {
		m[s] = true
	}
	for _, s := range b {
		_, ok := m[s]
		if ok {
			intersection = append(intersection, s)
		}
	}

	spew.Dump(intersection)
}
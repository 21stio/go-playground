package main

import (
	"regexp"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

func main() {
	sample := `*map[string]http.Request`

	var re = regexp.MustCompile(`([int|map|string|\[|\]|\*]+)|([a-zA-Z]+\.[a-zA-Z]+)|([a-zA-Z.]+)`)

	ms := re.FindAllSubmatch([]byte(sample), -1)

	types := []string{}
	for _, m := range ms {
		a := string(m[0])
		if a == string(m[3]) {
			a = "yo." + a
		}

		types = append(types, a)
	}

	spew.Dump(strings.Join(types, ""))

	//spew.Dump(ms)
}

package main

import (
	"unicode"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"github.com/davecgh/go-spew/spew"
	"fmt"
)

func main() {
	//s, err := normalize("12.000\u202f€")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//


	spew.Dump(fmt.Sprintf("%s", "12.000\u202f€"))
}

func normalize(s string) (s2 string, err error) {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
		return unicode.Is(unicode.Mn, r)
	}), norm.NFC)
	s2, _, err = transform.String(t, s)



	return
}

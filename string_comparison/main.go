package main

import (
	"unicode/utf8"
	"log"
	"bytes"
	"strings"
)

func compareStringsAsRunes() {
	string1 := "hello"
	string2 := "hello123"

	rune1, _ := utf8.DecodeRuneInString(string1)
	rune2, _ := utf8.DecodeRuneInString(string2)

	if (rune1 == rune2) {
		log.Print("compareStringsAsRunes is the same")
	} else {
		log.Print("compareStringsAsRunes is not the same")
	}
}

func compareStringsAsBytes() {
	string1 := "hello"
	string2 := "hello"

	bytes1 := []byte(string1)
	bytes2 := []byte(string2)

	if (bytes.Equal(bytes1, bytes2)) {
		log.Print("compareStringsAsBtyes is the same")
	} else {
		log.Print("compareStringsAsBtyes is not the same")
	}
}

func stringWithNewLineComparison()  {
	stringWithNewLine := "I am a string\nContaining new lines\n"
	stringWithOutNewLine := strings.Replace(stringWithNewLine,"\n", "" ,-1)

	if (strings.Compare(stringWithOutNewLine, "I am a stringContaining new lines") == 0) {
		log.Print("its the same, yo")
	}
}


func main() {
	stringWithNewLineComparison()
}

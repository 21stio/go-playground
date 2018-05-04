package main

import (
	"crypto/sha512"
	"io"
	"fmt"
	"encoding/hex"
)

func main() {

	//a := sha512.Sum512([]byte("asdsd"))
	//println(string(a[:]))
	//
	//sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	text := "His money is twice tainted: 'taint yours and 'taint mine."
	h := sha512.New()
	io.WriteString(h, text)
	fmt.Printf("% x", h.Sum(nil))

	algorithm := sha512.New()
	algorithm.Write([]byte(text))
	fmt.Printf(hex.EncodeToString(algorithm.Sum(nil)))

}

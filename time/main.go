package main

import (
	"time"
	"fmt"
)

func main() {
	ts, _ := time.Parse(time.RFC3339,"2018-05-31T22:08:41+00:00")

	since := time.Since(ts)

	fmt.Printf("%v", int(since.Hours()))
}

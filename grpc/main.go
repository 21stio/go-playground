package main

import (
	"github.com/davecgh/go-spew/spew"

	. "github.com/21stio/go-playground/grpc/schema/whatever"
	. "github.com/21stio/go-playground/grpc/schema/getwhateversresponse"
)

func main() {
	spew.Dump(Whatever{})
	spew.Dump(GetWhateversResponse{})
}

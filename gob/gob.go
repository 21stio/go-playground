package main

import (
	"net/http"
	"log"
	"encoding/gob"
	"bytes"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
)

type Response struct {
	Body []byte
	Header http.Header
}

func main() {
	err := func() (err error) {
		r, err := http.Get("http://ebay.com")
		if err != nil {
			return
		}

		resp := Response{}
		resp.Header = r.Header

		resp.Body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}

		gob.Register(Response{})

		var network bytes.Buffer
		enc := gob.NewEncoder(&network)
		dec := gob.NewDecoder(&network)
		err = enc.Encode(resp)
		if err != nil {
			return
		}

		var r2 Response
		err = dec.Decode(&r2)
		if err != nil {
			return
		}

		spew.Dump(r2)

		return
	}()
	if err != nil {
		log.Fatal(err)
	}
}

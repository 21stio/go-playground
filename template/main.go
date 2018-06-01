package main

import (
	"html/template"
	"os"
	"log"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {
	cwd, err := os.Getwd()
	if err != nil {
		return
	}

	tpml, err := template.ParseGlob(cwd + "/template/*.html")
	if err != nil {
		return
	}

	err = tpml.ExecuteTemplate(os.Stdout, "a.html", nil)
	if err != nil {
		return
	}

	return
}

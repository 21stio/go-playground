package main

import (
	"encoding/csv"
	"os"
	"log"
)

func main() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	f, err := os.Create("/tmp/dat2")
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}

	w := csv.NewWriter(f)
	w.WriteAll(records) // calls Flush internally

	records = [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}
	w.WriteAll(records)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

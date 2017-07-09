package main

import (
	"encoding/json"
	"log"
)

type Measurement struct {
	TemperatureCelsius float32
}

func main() {
	marshal()
}

func unmarshal() {
	j := "{\"temperature_celsius\": 0.5}"

	m := Measurement{}
	json.Unmarshal([]byte(j), &m)

	log.Printf("%+v", m)
}

func marshal()  {
	m := Measurement{
		TemperatureCelsius: 0.5,
	}

	b, _ := json.Marshal(&m)

	log.Print(string(b))
}

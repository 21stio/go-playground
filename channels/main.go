package main

import (
	"strconv"
	"time"
)


func main() {
	messageChan := make(chan string)
	consume(messageChan)

	time.Sleep(3*time.Second)

	produce(messageChan)
}

func produce(messages chan string) {
	println("start producing")

	names := []string{"Norberto Colombo","Andrea Grisby","Dreama Brwon","Mike Lobue","Tempie Borja","Vesta Wardlaw","Cristin Figueroa","Roselyn Varian","Fonda Au","Ivan Keese","Myrna Orsini","Remona Bethel","Chauncey Crichton","Michale Wheelwright","Lilly Nimmo","Euna Sesco","Oma Alfrey","Caprice Varughese","Sherice Mcaninch","Johnna Botsford","Lamont Kavanagh","Janeen Bonomo","Gracie Lovely","Huey Kitt","Anneliese Agan","Britney Brodbeck","Hallie Gosse","Hyo Ralston","Sharla Sloop","Patria Lamorte","Ammie Chapdelaine","Shantelle Hammell","Sharleen Reinert","Riley Levis","Misti Lukas","Vernetta Reese","Suzanne Hoekstra","Weldon Wotring","Mammie Zhou","Ranee Walker","Lean Miah","Dirk Labella","Cecilia Celaya","Charisse Husby","Zulma Navas","Samuel Krahn","Camellia Dennen","Lekisha Kerney","Debroah Friedt","Reva Feliciano"}

	for _, name := range names {
		messages <- name
	}
}

func consume(messages chan string){
	for i := 0; i < 10; i++ {
		go func(i int) {
			id := strconv.Itoa(i)
			println("started worker: " + id + ", waiting for messages")

			for message := range messages {
				time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				println("worker: " + string(id) + "   message: " + message)
			}
		}(i)
	}
}

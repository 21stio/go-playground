package main

import (
	"github.com/olebedev/emitter"
	"time"
)

func main() {
	e := &emitter.Emitter{}

	//go func() {
	//	for event := range e.On("change") {
	//		print(event.Int(0))
	//	}
	//}()

	go func() {
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		<-e.Emit("change", 37)
		print("yo")
	}()



	time.Sleep(10 * time.Second)
}

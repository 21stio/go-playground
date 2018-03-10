package main

import (
	"os"
	"time"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"log"
	"io/ioutil"
	"fmt"
)

func main()  {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {
	f, err := os.Open("beep/sound.wav")
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(f)
	fmt.Sprintf("%b", body)

	err = ioutil.WriteFile("beep/yo.txt", body, 644)


	s, format, err := wav.Decode(f)
	if err != nil {
		return
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan struct{})

	speaker.Play(beep.Seq(s, beep.Callback(func() {
		close(done)
	})))

	<-done

	return
}
package main

type LivingThing interface {
	Live()
}

type Lion struct {
	breathCount int
}

func (l *Lion) Live()  {
	l.breathCount++
}

func letItLive(t LivingThing)  {
	t.Live()
}

func main() {
	lion := Lion{}

	letItLive(lion)

	print(lion.breathCount)
}

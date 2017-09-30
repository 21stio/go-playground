package main

import (
	"go.uber.org/zap"
	"time"
	"errors"
)

func main() {
	l, _ := zap.NewProduction()
	sugar := l.Sugar()
	defer sugar.Sync()
	sugar.Infow("failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("failed to fetch URL: %s", "http://example.com")

	y := Yo{}
	y.Yo(sugar)
}

type Yo struct {

}

func (y *Yo) Yo(log *zap.SugaredLogger)  {
	log.Infof("aaaaa", "http://example.com")
	log.Error(errors.New("ouasd"))
}


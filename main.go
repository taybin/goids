package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"time"
)

var (
	count = kingpin.Flag("count", "Number of boids").Short('c').Default("500").Uint()
)

type appContext struct {
	area *Area
}

func main() {
	kingpin.Version("0.7").Author("Taybin Rutkin")
	kingpin.CommandLine.Help = "An implementation of boids"
	kingpin.Parse()
	context := &appContext{
		area: NewArea(300, 300, 300),
	}

	setupBoids(context)
	setupWeb(context)
}

func setupBoids(context *appContext) {
	var lastID = 0

	for i := 0; i < int(*count); i++ {
		lastID = lastID + 1
		boid := NewBoid(lastID)
		boid.RandomizePosition(context.area)
		context.area.AddBoid(boid)
	}

	go func() {
		for {
			context.area.UpdateBoids()
			time.Sleep(time.Second * 4)
		}
	}()
}

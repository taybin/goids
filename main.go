package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"time"
)

var (
	count = kingpin.Flag("count", "Number of boids").Short('c').Default("500").Uint()
	xD    = Dimension(kingpin.Arg("x-dimension", "Start and stop of x dimension").Default("0:100"))
	yD    = Dimension(kingpin.Arg("y-dimension", "Start and stop of y dimension").Default("0:100"))
	zD    = Dimension(kingpin.Arg("z-dimension", "Start and stop of z dimension").Default("0:100"))
)

type appContext struct {
	area *Area
}

func main() {
	kingpin.Version("0.7").Author("Taybin Rutkin")
	kingpin.CommandLine.Help = "An implementation of boids"
	kingpin.Parse()
	log.Printf("x-dimension %f - %f\n", xD.Start, xD.Stop)
	log.Printf("y-dimension %f - %f\n", yD.Start, yD.Stop)
	log.Printf("z-dimension %f - %f\n", zD.Start, zD.Stop)

	context := &appContext{
		area: NewArea(xD, yD, zD),
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

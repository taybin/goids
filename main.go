package main

import (
	"time"
)

type appContext struct {
	area *Area
}

func main() {
	context := &appContext{
		area: NewArea(300, 300, 300, 98, 98, 98),
	}

	setupBoids(context)
	setupWeb(context)
}

func setupBoids(context *appContext) {
	var lastID = 0

	for i := 0; i < 500; i++ {
		lastID = lastID + 1
		boid := NewBoid(lastID, len(context.area.Dimensions))
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

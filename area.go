package main

import (
	rtree "github.com/dhconnelly/rtreego"
)

type Area struct {
	Tree       *rtree.Rtree
	Dimensions []int32
	Boids      map[int]*Boid
	SendChan   chan *Boid
}

func NewArea(dimensions ...int32) *Area {
	n_dimensions := len(dimensions)

	a := &Area{
		Tree:       rtree.NewTree(n_dimensions, 25, 50),
		Dimensions: dimensions,
		Boids:      make(map[int]*Boid),
		SendChan:   make(chan *Boid),
	}
	return a
}

func (a *Area) AddBoid(boid *Boid) {
	a.Boids[boid.ID] = boid
	a.Tree.Insert(boid)
}

func (a *Area) UpdateBoids() {
	for _, boid := range a.Boids {
		boid.UpdateVelocity(a)
		boid.UpdatePosition()
		a.SendChan <- boid
		a.Tree.Delete(boid)
		a.Tree.Insert(boid)
	}
}

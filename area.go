package main

import (
	rtree "github.com/dhconnelly/rtreego"
)

type Area struct {
	Tree       *rtree.Rtree
	Dimensions []int32
	Boids      map[int]*Boid
	LastID     int
	SendChan   chan *Boid
}

func NewArea(dimensions ...int32) *Area {
	n_dimensions := len(dimensions)

	a := &Area{
		Tree:       rtree.NewTree(n_dimensions, 25, 50),
		Dimensions: dimensions,
		Boids:      make(map[int]*Boid),
		SendChan:   make(chan *Boid),
		LastID:     0,
	}
	return a
}

func (a *Area) AddBoid() {
	id := a.LastID + 1
	boid := NewBoid(id, a.Dimensions)
	a.Boids[id] = boid
	a.Tree.Insert(boid)
	a.LastID = id
}

func (a *Area) UpdateBoids() {
	for _, boid := range a.Boids {
		boid.UpdateVelocity(a)
	}
	for _, boid := range a.Boids {
		boid.UpdatePosition()
		a.SendChan <- boid
	}
}

package main

import (
	rtree "github.com/patrick-higgins/rtreego"
)

type Area struct {
	Tree     *rtree.Rtree
	X        *DimensionValue
	Y        *DimensionValue
	Z        *DimensionValue
	Boids    map[int]*Boid
	SendChan chan *Boid
}

func NewArea(xD, yD, zD *DimensionValue) *Area {
	a := &Area{
		Tree:     rtree.NewTree(25, 50),
		X:        xD,
		Y:        yD,
		Z:        zD,
		Boids:    make(map[int]*Boid),
		SendChan: make(chan *Boid),
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

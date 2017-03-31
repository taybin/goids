// With many thanks to http://www.kfish.org/boids/pseudocode.html

package main

import (
	rtree "github.com/dhconnelly/rtreego"
	"math/rand"
)

type Boid struct {
	ID       int
	Point    rtree.Point
	Velocity []float64
}

func NewBoid(id int, dimensions []int32) *Boid {
	var point rtree.Point
	var vel []float64

	for i := range dimensions {
		p := rand.Int31n(dimensions[i])
		point = append(point, float64(p))
		vel = append(vel, 0.0)
	}

	return &Boid{
		ID:       id,
		Point:    point,
		Velocity: vel,
	}
}

func (b *Boid) Bounds() *rtree.Rect {
	return b.Point.ToRect(0.01)
}

func (b *Boid) UpdateVelocity(area *Area) {
	v1 := b.rule1(area)
	v2 := b.rule2(area)
	v3 := b.rule3(area)

	b.Velocity = addFloats(v1, v2, v3)
}

func (b *Boid) UpdatePosition() {
	b.Point = addFloats(b.Point, b.Velocity)
}

// rule1
//
// 	PROCEDURE rule1(boid bJ)
//
//		Vector pcJ
//
//    FOR EACH BOID b
//    	IF b != bJ THEN
//    		pcJ = pcJ + b.position
//    	END IF
//    END
//
//    pcJ = pcJ / N-1
//
//    RETURN (pcJ - bJ.position) / 100
//
// 	END PROCEDURE
func (b *Boid) rule1(area *Area) []float64 {
	pcJ := make(rtree.Point, len(b.Point))
	for i := range pcJ {
		pcJ[i] = 0
	}

	for id, boid := range area.Boids {
		if id != b.ID {
			for k, v := range boid.Point {
				pcJ[k] = pcJ[k] + v
			}
		}
	}

	for i := range pcJ {
		pcJ[i] = pcJ[i] / float64(len(area.Boids)-1)
	}

	subbed := SubFloats(b.Point, pcJ)
	divved := divFloat(subbed, 100.0)
	return divved
}

func (b *Boid) rule2(area *Area) []float64 {
	return make([]float64, 1)
}

func (b *Boid) rule3(area *Area) []float64 {
	return make([]float64, 1)
}

func SubFloats(a, b []float64) []float64 {
	newPoint := make([]float64, len(a))
	for i := range a {
		newPoint[i] = a[i] - b[i]
	}
	return newPoint
}

func divFloat(a []float64, div float64) []float64 {
	newPoint := make([]float64, len(a))
	for i := range a {
		newPoint[i] = a[i] / div
	}
	return newPoint
}

func addFloats(floats ...[]float64) []float64 {
	newFloat := makeFloats(int32(len(floats[0])))

	for i := range floats {
		for j := range floats[i] {
			newFloat[j] = newFloat[j] + floats[i][j]
		}
	}

	return newFloat
}

func makeFloats(size int32) []float64 {
	newFloats := make([]float64, size)

	for i := range newFloats {
		newFloats[i] = 0
	}

	return newFloats
}

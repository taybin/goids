// With many thanks to http://www.kfish.org/boids/pseudocode.html

package main

import (
	rtree "github.com/dhconnelly/rtreego"
	// "log"
	"math"
	"math/rand"
)

type Boid struct {
	ID       int         `json:"id"`
	Point    rtree.Point `json:"position"`
	Velocity []float64   `json:"-"`
}

func NewBoid(id, dimensions int) *Boid {
	var point rtree.Point
	var vel []float64

	for i := 0; i < dimensions; i++ {
		point = append(point, 0.0)
		vel = append(vel, 0.0)
	}

	return &Boid{
		ID:       id,
		Point:    point,
		Velocity: vel,
	}
}

func (b *Boid) RandomizePosition(area *Area) {
	var point rtree.Point

	for i := range area.Dimensions {
		p := rand.Int31n(area.Dimensions[i])
		point = append(point, float64(p))
	}

	b.Point = point
}

func (b *Boid) Bounds() *rtree.Rect {
	return b.Point.ToRect(0.01)
}

func (b *Boid) UpdateVelocity(area *Area) {
	v1 := b.Rule1(area)
	v2 := b.Rule2(area)
	v3 := b.Rule3(area)

	b.Velocity = AddFloats(v1, v2, v3)
}

func (b *Boid) UpdatePosition() {
	b.Point = AddFloats(b.Point, b.Velocity)
}

// Rule1
//
// 	PROCEDURE Rule1(boid bJ)
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
func (b *Boid) Rule1(area *Area) []float64 {
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

	subbed := SubFloats(pcJ, b.Point)
	divved := DivFloat(subbed, 100.0)
	return divved
}

// Rule2
//
// PROCEDURE rule2(boid bJ)
//
// 	Vector c = 0;
//
// 	FOR EACH BOID b
// 		IF b != bJ THEN
// 			IF |b.position - bJ.position| < 100 THEN
// 				c = c - (b.position - bJ.position)
// 			END IF
// 		END IF
// 	END
//
// 	RETURN c
//
// END PROCEDURE
func (b *Boid) Rule2(area *Area) []float64 {
	vector := makeFloats(int32(len(b.Point)))

	for id, boid := range area.Boids {
		if id != b.ID {
			for k, v := range boid.Point {
				if math.Abs(v-b.Point[k]) < 10 {
					vector[k] = vector[k] - (v - b.Point[k])
				}
			}
		}
	}

	return vector
}

// Rule3
//
// PROCEDURE rule3(boid bJ)
//
// 	Vector pvJ
//
// 	FOR EACH BOID b
// 		IF b != bJ THEN
// 			pvJ = pvJ + b.velocity
// 		END IF
// 	END
//
// 	pvJ = pvJ / N-1
//
// 	RETURN (pvJ - bJ.velocity) / 8
//
// END PROCEDURE
func (b *Boid) Rule3(area *Area) []float64 {
	pvJ := makeFloats(int32(len(b.Velocity)))

	for id, boid := range area.Boids {
		if id != b.ID {
			for k, v := range boid.Velocity {
				pvJ[k] = pvJ[k] + v
			}
		}
	}

	for i := range pvJ {
		pvJ[i] = pvJ[i] / float64(len(area.Boids)-1)
	}

	subbed := SubFloats(pvJ, b.Velocity)
	divved := DivFloat(subbed, 8.0)
	return divved
}

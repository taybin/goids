// With many thanks to http://www.kfish.org/boids/pseudocode.html

package main

import (
	rtree "github.com/patrick-higgins/rtreego"
	"math"
)

type Boid struct {
	ID       int         `json:"id"`
	Point    rtree.Point `json:"position"`
	Velocity rtree.Point `json:"-"`
}

func NewBoid(id int) *Boid {
	return &Boid{
		ID:       id,
		Point:    rtree.Point{0.0, 0.0, 0.0},
		Velocity: rtree.Point{0.0, 0.0, 0.0},
	}
}

func (b *Boid) RandomizePosition(area *Area) {
	b.Point[0] = area.X.Rand()
	b.Point[1] = area.Y.Rand()
	b.Point[2] = area.Z.Rand()
}

func (b *Boid) Bounds() *rtree.Rect {
	return b.Point.ToRect(0.01)
}

func (b *Boid) UpdateVelocity(area *Area) {
	v1 := b.Rule1(area)
	v2 := b.Rule2(area)
	v3 := b.Rule3(area)

	velocities := AddPoints(b.Velocity, v1, v2, v3)
	b.Velocity = LimitVelocity(velocities)
}

func (b *Boid) UpdatePosition() {
	b.Point = AddPoints(b.Point, b.Velocity)
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
func (b *Boid) Rule1(area *Area) rtree.Point {
	pcJ := rtree.Point{0.0, 0.0, 0.0}

	nearest := area.Tree.NearestNeighbors(5, b.Point)
	boids := SpatialsToBoids(nearest)

	for _, boid := range boids {
		if boid.ID != b.ID {
			for i := 0; i < 3; i++ {
				pcJ[i] = pcJ[i] + boid.Point[i]
			}
		}
	}

	for i := 0; i < 3; i++ {
		pcJ[i] = pcJ[i] / float64(len(boids)-1)
	}

	subbed := SubPoints(pcJ, b.Point)
	divved := DivPoint(subbed, 100.0)
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
func (b *Boid) Rule2(area *Area) rtree.Point {
	vector := rtree.Point{0.0, 0.0, 0.0}

	nearest := area.Tree.NearestNeighbors(5, b.Point)
	boids := SpatialsToBoids(nearest)

	for _, boid := range boids {
		if boid.ID != b.ID {
			for i := 0; i < 3; i++ {
				if math.Abs(boid.Point[i]-b.Point[i]) < 10 {
					vector[i] = vector[i] - (boid.Point[i] - b.Point[i])
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
func (b *Boid) Rule3(area *Area) rtree.Point {
	pvJ := rtree.Point{0.0, 0.0, 0.0}

	nearest := area.Tree.NearestNeighbors(5, b.Point)
	boids := SpatialsToBoids(nearest)

	for _, boid := range boids {
		if boid.ID != b.ID {
			for i := 0; i < 3; i++ {
				pvJ[i] = pvJ[i] + boid.Velocity[i]
			}
		}
	}

	for i := 0; i < 3; i++ {
		pvJ[i] = pvJ[i] / float64(len(boids)-1)
	}

	subbed := SubPoints(pvJ, b.Velocity)
	divved := DivPoint(subbed, 8.0)
	return divved
}

// LimitVelocity
//
//  PROCEDURE limit_velocity(Boid b)
//          Integer vlim
//          Vector v
//
//          IF |b.velocity| > vlim THEN
//                  b.velocity = (b.velocity / |b.velocity|) * vlim
//          END IF
//  END PROCEDURE
func LimitVelocity(velocities rtree.Point) rtree.Point {
	vlim := 25.0
	var absVel float64

	for i := 0; i < 3; i++ {
		absVel = math.Abs(velocities[i])
		if absVel > vlim {
			velocities[i] = (velocities[i] / absVel) * vlim
		}
	}

	return velocities
}

func SpatialsToBoids(spatials []rtree.Spatial) []*Boid {
	var boids []*Boid

	for _, spatial := range spatials {
		if spatial != nil {
			boids = append(boids, spatial.(*Boid))
		}
	}
	return boids
}

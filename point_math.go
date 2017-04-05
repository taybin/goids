package main

import (
	rtree "github.com/patrick-higgins/rtreego"
)

// SubPoints subtracts one rtree.Point from another
// and returns a new rtree.Point with the result
func SubPoints(a, b rtree.Point) rtree.Point {
	newPoint := rtree.Point{0.0, 0.0, 0.0}

	for i := 0; i < 3; i++ {
		newPoint[i] = a[i] - b[i]
	}
	return newPoint
}

// DivPoint divides each value in an rtree.Point by a
// float64 and returns a new rtree.Point with the result
func DivPoint(a rtree.Point, div float64) rtree.Point {
	newPoint := rtree.Point{0.0, 0.0, 0.0}

	for i := 0; i < 3; i++ {
		newPoint[i] = a[i] / div
	}
	return newPoint
}

// AddPoints accepts a varadic amount of rtree.Points,
// adds them all together, and returns a new rtree.Point with
// the result
func AddPoints(points ...rtree.Point) rtree.Point {
	newPoint := rtree.Point{0.0, 0.0, 0.0}

	for i := range points {
		for j := 0; j < 3; j++ {
			newPoint[j] = newPoint[j] + points[i][j]
		}
	}

	return newPoint
}

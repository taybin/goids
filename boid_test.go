package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
)

var _ = Describe("Boids", func() {
	Describe("Boid Initialization", func() {
		var (
			boid *Boid
		)
		BeforeEach(func() {
			boid = NewBoid(1, 3)
		})
		It("Should proper dimensions", func() {
			Expect(boid.Point).To(HaveLen(3))
		})
		It("Should have zero'd position", func() {
			Expect(boid.Point).To(ConsistOf([]float64{0.0, 0.0, 0.0}))
		})
		It("Should have zero'd velocity", func() {
			Expect(boid.Velocity).To(ConsistOf([]float64{0.0, 0.0, 0.0}))
		})
	})
	Describe("Single Dimension Tests", func() {
		var (
			area  *Area
			boid1 *Boid
			boid2 *Boid
		)
		BeforeEach(func() {
			area = NewArea(100)
			boid1 = NewBoid(1, 1)
			boid1.Point = []float64{25.0}
			boid2 = NewBoid(2, 1)
			boid2.Point = []float64{75.0}
			area.AddBoid(boid1)
			area.AddBoid(boid2)
		})

		It("Should implement rule1", func() {
			result := boid1.Rule1(area)
			log.Println(result)
			// Expect(result)
		})
	})
})

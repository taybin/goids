package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
	Describe("Rule tests", func() {
		var (
			area  *Area
			boid1 *Boid
			boid2 *Boid
		)
		JustBeforeEach(func() {
			area.AddBoid(boid1)
			area.AddBoid(boid2)
		})
		Context("Single Dimension Tests", func() {
			BeforeEach(func() {
				area = NewArea(100)
				boid1 = NewBoid(1, 1)
				boid1.Point = []float64{25.0}
				boid2 = NewBoid(2, 1)
				boid2.Point = []float64{75.0}
			})

			It("Should implement rule1", func() {
				result := boid1.Rule1(area)
				Expect(result).To(ConsistOf([]float64{0.5}))
			})

			It("Should implement rule2", func() {
				result := boid1.Rule2(area)
				Expect(result).To(ConsistOf([]float64{0.0}))
			})
		})

		Context("Two Dimension Tests", func() {
			BeforeEach(func() {
				area = NewArea(100, 100)
				boid1 = NewBoid(1, 2)
				boid1.Point = []float64{25.0, 50.0}
				boid2 = NewBoid(2, 2)
				boid2.Point = []float64{75.0, 25.0}
			})

			It("Should implement rule1", func() {
				result := boid1.Rule1(area)
				Expect(result).To(ConsistOf([]float64{0.5, -0.25}))
			})

			It("Should implement rule2", func() {
				result := boid1.Rule2(area)
				Expect(result).To(ConsistOf([]float64{0.0, 0.0}))
			})
		})
	})
})

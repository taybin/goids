package main

import (
	rtree "github.com/patrick-higgins/rtreego"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Boids", func() {
	Describe("Boid Initialization", func() {
		var (
			boid *Boid
		)
		BeforeEach(func() {
			boid = NewBoid(1)
		})
		It("Should proper dimensions", func() {
			Expect(boid.Point).To(HaveLen(3))
		})
		It("Should have zero'd position", func() {
			Expect(boid.Point).To(ConsistOf(rtree.Point{0.0, 0.0, 0.0}))
		})
		It("Should have zero'd velocity", func() {
			Expect(boid.Velocity).To(ConsistOf(rtree.Point{0.0, 0.0, 0.0}))
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
				boid1 = NewBoid(1)
				boid1.Point = rtree.Point{45.0, 0.0, 0.0}
				boid2 = NewBoid(2)
				boid2.Point = rtree.Point{55.0, 0.0, 0.0}
			})

			It("Should implement rule1", func() {
				result := boid1.Rule1(area)
				Expect(result).To(ConsistOf(rtree.Point{0.1, 0.0, 0.0}))
			})

			It("Should implement rule2", func() {
				result := boid1.Rule2(area)
				Expect(result).To(ConsistOf(rtree.Point{0.0, 0.0, 0.0}))
			})
		})

		Context("Two Dimension Tests", func() {
			BeforeEach(func() {
				area = NewArea(100, 100)
				boid1 = NewBoid(1)
				boid1.Point = rtree.Point{25.0, 50.0, 0.0}
				boid2 = NewBoid(2)
				boid2.Point = rtree.Point{75.0, 25.0, 0.0}
			})

			It("Should implement rule1", func() {
				result := boid1.Rule1(area)
				Expect(result).To(ConsistOf(rtree.Point{0.5, -0.25, 0.0}))
			})

			It("Should implement rule2", func() {
				result := boid1.Rule2(area)
				Expect(result).To(ConsistOf(rtree.Point{0.0, 0.0, 0.0}))
			})
		})
	})
})

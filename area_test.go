package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Area", func() {
	Describe("Area Initialization", func() {
		var (
			area *Area
		)
		BeforeEach(func() {
			xD := NewDimensionValue(0.0, 100.0)
			yD := NewDimensionValue(0.0, 100.0)
			zD := NewDimensionValue(0.0, 100.0)
			area = NewArea(xD, yD, zD)
		})
		It("Should have zero boids", func() {
			Expect(area.Boids).To(HaveLen(0))
		})
	})
})

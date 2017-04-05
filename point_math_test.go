package main

import (
	rtree "github.com/patrick-higgins/rtreego"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Point Math", func() {
	Describe("SubPoints", func() {
		It("Should subtract points", func() {
			a := rtree.Point{2.0, 0.0, 0.0}
			b := rtree.Point{1.0, 0.0, 0.0}
			result := SubPoints(a, b)
			Expect(result).To(ConsistOf(rtree.Point{1.0, 0.0, 0.0}))
		})
		It("Should subtract multi-points", func() {
			a := rtree.Point{2.0, 4.0}
			b := rtree.Point{1.0, 2.0}
			result := SubPoints(a, b)
			Expect(result).To(ConsistOf(rtree.Point{1.0, 2.0}))
		})
	})

	Describe("DivPoint", func() {
		It("Should divide points", func() {
			a := rtree.Point{100.0, 0.0, 0.0}
			result := DivPoint(a, 100)
			Expect(result).To(ConsistOf(rtree.Point{1.0, 0.0, 0.0}))
		})
		It("Should divide multi-points", func() {
			a := rtree.Point{100.0, 200.0, 0.0}
			result := DivPoint(a, 100)
			Expect(result).To(ConsistOf(rtree.Point{1.0, 2.0, 0.0}))
		})
	})

	Describe("AddPoints", func() {
		It("Should add points", func() {
			a := rtree.Point{10.0, 0.0, 0.0}
			b := rtree.Point{15.0, 0.0, 0.0}
			result := AddPoints(a, b)
			Expect(result).To(ConsistOf(rtree.Point{25.0, 0.0, 0.0}))
		})
		It("Should add multi-points", func() {
			a := rtree.Point{10.0, 5.5, 0.0}
			b := rtree.Point{15.0, 7.25, 0.0}
			result := AddPoints(a, b)
			Expect(result).To(ConsistOf(rtree.Point{25.0, 12.75, 0.0}))
		})
	})
})

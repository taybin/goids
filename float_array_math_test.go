package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Float Array Math", func() {
	Describe("SubFloats", func() {
		It("Should subtract floats", func() {
			a := []float64{2.0}
			b := []float64{1.0}
			result := SubFloats(a, b)
			Expect(result).To(ConsistOf([]float64{1.0}))
		})
		It("Should subtract multi-floats", func() {
			a := []float64{2.0, 4.0}
			b := []float64{1.0, 2.0}
			result := SubFloats(a, b)
			Expect(result).To(ConsistOf([]float64{1.0, 2.0}))
		})
	})

	Describe("DivFloat", func() {
		It("Should divide floats", func() {
			a := []float64{100.0}
			result := DivFloat(a, 100)
			Expect(result).To(ConsistOf([]float64{1.0}))
		})
		It("Should divide multi-floats", func() {
			a := []float64{100.0, 200.0}
			result := DivFloat(a, 100)
			Expect(result).To(ConsistOf([]float64{1.0, 2.0}))
		})
	})

	Describe("AddFloats", func() {
		It("Should add floats", func() {
			a := []float64{10.0}
			b := []float64{15.0}
			result := AddFloats(a, b)
			Expect(result).To(ConsistOf([]float64{25.0}))
		})
		It("Should add multi-floats", func() {
			a := []float64{10.0, 5.5}
			b := []float64{15.0, 7.25}
			result := AddFloats(a, b)
			Expect(result).To(ConsistOf([]float64{25.0, 12.75}))
		})
	})

	Describe("makeFloats", func() {
		It("Should return zero'd float array", func() {
			result := makeFloats(3)
			Expect(result).To(HaveLen(3))
			Expect(result).To(ConsistOf([]float64{0.0, 0.0, 0.0}))
		})
	})
})

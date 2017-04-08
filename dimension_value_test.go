package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DimensionValue", func() {
	It("Should initialize with floats", func() {
		dv := NewDimensionValue(10.0, 20.0)
		Expect(dv.Start).To(Equal(10.0))
		Expect(dv.Stop).To(Equal(20.0))
	})

	It("Should initialize with inits", func() {
		dv := NewDimensionValue(10, 20)
		Expect(dv.Start).To(Equal(10.0))
		Expect(dv.Stop).To(Equal(20.0))
	})

	It("Should parse a string", func() {
		dv := &DimensionValue{}
		dv.Set("-1:30")
		Expect(dv.Start).To(Equal(-1.0))
		Expect(dv.Stop).To(Equal(30.0))
	})

	It("Should return a bounded, random value", func() {
		dvs := []*DimensionValue{
			{0, 1},
			{-50, 50},
			{100, 200},
			{0, 0},
			{0.0, 0.0},
			{1.0, 1.0},
		}
		for _, dv := range dvs {
			r := dv.Rand()
			Expect(r).To(BeNumerically(">=", dv.Start))
			Expect(r).To(BeNumerically("<=", dv.Stop))
		}
	})
})

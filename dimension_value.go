package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"math/rand"
	"strconv"
	"strings"
)

func Dimension(s kingpin.Settings) *DimensionValue {
	value := &DimensionValue{}
	s.SetValue(value)
	return value
}

type DimensionValue struct {
	Start float64
	Stop  float64
}

func NewDimensionValue(start, stop float64) *DimensionValue {
	return &DimensionValue{
		Start: start,
		Stop:  stop,
	}
}

func (dv *DimensionValue) Set(value string) error {
	parts := strings.SplitN(value, ":", 2)
	if len(parts) != 2 {
		return fmt.Errorf("expected START:STOP got '%s'", value)
	}

	var f64 float64
	var err error

	f64, err = strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return err
	}
	dv.Start = f64

	f64, err = strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return err
	}
	dv.Stop = f64

	return nil
}

func (dv *DimensionValue) String() string {
	return ""
}

// Rand() returns a bounded value between or equal to the Start (inclusive)
// and Stop (exclusive) values.  It returns [Start,Stop)
func (dv *DimensionValue) Rand() float64 {
	return ((dv.Stop - dv.Start) * rand.Float64()) + dv.Start
}

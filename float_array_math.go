package main

func SubFloats(a, b []float64) []float64 {
	newPoint := make([]float64, len(a))
	for i := range a {
		newPoint[i] = a[i] - b[i]
	}
	return newPoint
}

func DivFloat(a []float64, div float64) []float64 {
	newPoint := make([]float64, len(a))
	for i := range a {
		newPoint[i] = a[i] / div
	}
	return newPoint
}

func AddFloats(floats ...[]float64) []float64 {
	newFloat := makeFloats(int32(len(floats[0])))

	for i := range floats {
		for j := range floats[i] {
			newFloat[j] = newFloat[j] + floats[i][j]
		}
	}

	return newFloat
}

func makeFloats(size int32) []float64 {
	newFloats := make([]float64, size)

	for i := range newFloats {
		newFloats[i] = 0
	}

	return newFloats
}

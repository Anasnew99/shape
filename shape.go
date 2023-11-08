package shape

type Shape interface {
	Area() float64
	Perimeter() float64
}

func CalculateBoundaries(s []Shape) float64 {
	var boundaries float64
	for _, shape := range s {
		boundaries += shape.Perimeter()
	}

	return boundaries
}

func CalculateTotalSurfaceArea(s []Shape) float64 {
	var area float64
	for _, shape := range s {
		area += shape.Area()
	}

	return area
}

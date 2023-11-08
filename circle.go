package shape

type circle struct {
	radius float64
}

func NewCircle(radius float64) circle {
	r := circle{
		radius,
	}
	return r
}

func (r circle) Area() float64 {
	return 3.14 * r.radius * r.radius
}

func (r circle) Perimeter() float64 {
	return 2 * 3.14 * r.radius
}

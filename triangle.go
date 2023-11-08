package shape

type triangle struct {
	a float64
	b float64
	c float64
}

func NewTriangle(a float64, b float64, c float64) triangle {
	r := triangle{
		a,
		b,
		c,
	}
	return r
}

func (r triangle) Area() float64 {
	return 0.5 * r.a * r.b
}

func (r triangle) Perimeter() float64 {
	return r.a + r.b + r.c
}

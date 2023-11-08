package shape

type rectangle struct {
	width  float64
	height float64
}

func NewRectangle(width float64, height float64) rectangle {
	r := rectangle{
		width,
		height,
	}
	return r
}

func (r rectangle) Area() float64 {
	return r.width * r.height
}

func (r rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

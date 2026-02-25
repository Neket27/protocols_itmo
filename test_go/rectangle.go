package test_go

type Rectangle struct {
	width  float64
	height float64
}

func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{width, height}
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

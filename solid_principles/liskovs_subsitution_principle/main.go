package main

import "fmt"

//liskov substitution principle

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func UserIt(sized Sized) {

	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	acrualArea := sized.GetHeight() * sized.GetWidth()
	fmt.Println(expectedArea, acrualArea)
}

func main() {

	rc := &Rectangle{5, 3}
	fmt.Println(rc)
	UserIt(rc)

	sq := NewSquare(5)
	UserIt(sq)
}

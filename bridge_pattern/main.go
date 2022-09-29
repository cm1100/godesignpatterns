package main

import "fmt"

// Circle Square
// Raster, Vector

// RasterCircle , VectorCircle , RasterSquare,

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRendrer struct {
}

func (v *VectorRendrer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius ", radius)
}

type RasterRenderer struct {
	Dpi int
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	//TODO implement me
	fmt.Println("Drawing pixels for a circle of radius ", radius)
}

type Circle struct {
	render Renderer
	radius float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer, radius}
}

func (c *Circle) Draw() {
	c.render.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor

}

func main() {
	raster := RasterRenderer{}
	vector := VectorRendrer{}
	circle := NewCircle(&raster, 5)
	circle2 := NewCircle(&vector, 10)
	circle.Draw()
	circle.Resize(2)
	circle2.Draw()
	circle2.Resize(5)
	circle2.Draw()
}

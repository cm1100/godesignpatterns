package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of Radius %f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with size %f", s.Side)
}

type ColoredShape struct {
	shape Shape
	Color string
}

func (c *ColoredShape) Render() string {

	return fmt.Sprintf("%s has the color %s", c.shape.Render(), c.Color)

}

type TransparentShape struct {
	shape        Shape
	transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has a %f transparency ", t.shape.Render(), t.transparency)
}

func main() {

	circle := Circle{2}
	fmt.Println(circle.Render())

	redCircle := ColoredShape{&circle, "black"}
	fmt.Println(redCircle.Render())

	redCircletransparent := TransparentShape{&redCircle, 0.90}
	fmt.Println(redCircletransparent.Render())
}

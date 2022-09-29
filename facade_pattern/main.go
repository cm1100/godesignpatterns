package main

import "fmt"

// Build Facade over internals

type Console struct {
	buffer    []*Buffer
	viewports []*ViewPort
	offset    int
}

func NewConsole() *Console {
	b := NewBuffer(200, 150)
	v := NewViewPort(b)
	return &Console{[]*Buffer{b}, []*ViewPort{v}, 10}
}

func (c *Console) GetCharacterAt(index int) rune {
	return c.viewports[0].GetCharacterAt(index)
}

func main() {

	c := NewConsole()
	u := c.GetCharacterAt(10)
	fmt.Println(u)
}

package main

import "fmt"

type Aged interface {
	Age() int
	SetAge(age int)
}

type Bird struct {
	age int
}

func (b *Bird) Age() int {
	return b.age
}

func (b *Bird) SetAge(age int) {
	b.age = age
}

func (b *Bird) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying")
	}
}

type Lizard struct {
	age int
}

func (l *Lizard) Age() int {
	return l.age
}

func (l *Lizard) SetAge(age int) {
	l.age = age
}

type Dragon struct {
	bird    Bird
	lizzard Lizard
}

func (d *Dragon) Age() int {
	return d.bird.age
}

func (d *Dragon) SetAge(age int) {
	d.bird.SetAge(age)
	d.lizzard.SetAge(age)
}

func (d *Dragon) Fly() {
	d.bird.Fly()
}

func NewDragon() *Dragon {
	return &Dragon{Bird{}, Lizard{}}
}

func main() {

	d := Dragon{}
	d.SetAge(1)
	d.Fly()

}

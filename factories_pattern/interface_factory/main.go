package main

import "fmt"

type Person interface {
	sayHello()
}

type person struct {
	name string
	age  int
}

func (p *person) sayHello() {
	fmt.Println("Hi my name is ", p.name)
}

type tiredPerson struct {
	name string
	age  int
}

func (p *tiredPerson) sayHello() {
	fmt.Println("well done")
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{
			name, age,
		}
	}

	return &person{name: name, age: age}

}

func main() {

	p := NewPerson("kames", 22)
	p.sayHello()
	p2 := NewPerson("ko", 800)
	p2.sayHello()
}

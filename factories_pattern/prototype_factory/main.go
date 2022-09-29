package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "developer", 60000}
	case Manager:
		return &Employee{"", "manager", 100000}
	default:
		panic("Unsupported role")
	}

}

func main() {

	m := NewEmployee(Manager)
	m.Name = "aca"
	fmt.Println(m)
	dev := NewEmployee(Developer)
	fmt.Println(dev)
}

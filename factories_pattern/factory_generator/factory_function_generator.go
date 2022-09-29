package main

import "fmt"

type Employee struct {
	Name, position string
	AnnualIncome   int
}

// functional approach
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {

	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}

}

func UseFunctionGenartor() {
	developerFactory := NewEmployeeFactory("developer", 60000)
	managerFactory := NewEmployeeFactory("manager", 100000)
	developer1 := developerFactory("chaitanya")
	manager1 := managerFactory("aca")
	fmt.Println(developer1)
	fmt.Println(manager1)
}

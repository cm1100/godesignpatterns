package main

import "fmt"

type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func NewEmployeeFactoryStruct(position string, annualincome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualincome}
}

func UseStruct() {

	boss := NewEmployeeFactoryStruct("ceo", 100000)
	sus := boss.Create("sush")
	fmt.Println(sus)
}

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Address
}

func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)
	d := gob.NewDecoder(&b)
	result := Employee{}
	_ = d.Decode(&result)
	return &result
}

var manOffice = Employee{
	"", Address{0, "123 east drive", "London"},
}

var auxOffice = Employee{
	"", Address{0, "66 west drive drive", "London"},
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {

	return newEmployee(&manOffice, name, suite)

}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func main() {

	john := NewMainOfficeEmployee("john", 100)
	jane := NewAuxOfficeEmployee("Jane", 200)
	fmt.Println(john, jane)
}

package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {

	return &Address{
		a.StreetAddress,
		a.City,
		a.Country,
	}

}

type Person struct {
	name    string
	address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.address = p.address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q

}

// deep copying , making copy of all the objects , pointers , slices and much more if you dont do this changes conflict with one another

func main() {
	john := Person{
		"john",
		&Address{"13 London", "London", "Uk"},
		[]string{"chaitanya1", "chaitanya 2"},
	}

	jane3 := john.DeepCopy()
	fmt.Println(jane3)
}

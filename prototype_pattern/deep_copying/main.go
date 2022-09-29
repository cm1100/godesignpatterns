package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	name    string
	address *Address
}

// deep copying , making copy of all the objects , pointers , slices and much more if you dont do this changes conflict with one another

func main() {
	john := Person{
		"john",
		&Address{"13 London", "London", "Uk"},
	}
	jane := john
	jane.name = "jane"
	jane.address.StreetAddress = "play"
	fmt.Println(jane.address)
	fmt.Println(john.address)

	jane2 := john
	jane2.address = &Address{
		john.address.StreetAddress,
		john.address.City,
		"India",
	}
	fmt.Println(john.address)
	fmt.Println(jane2.address)
}

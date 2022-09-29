package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	name    string
	address *Address
	Friends []string
}

// deep copying , making copy of all the objects , pointers , slices and much more if you dont do this changes conflict with one another

func (p *Person) DeepCopy() *Person {
	// making memory
	b := bytes.Buffer{}
	// giving encoder memory to write to
	e := gob.NewEncoder(&b)
	// added person using encoder to memory
	_ = e.Encode(p)
	fmt.Println(string(b.Bytes()))

	// giving memory for reading of decoder
	d := gob.NewDecoder(&b)

	result := Person{}
	// reading memory into person object
	_ = d.Decode(&result)

	return &result
}

func main() {
	john := Person{
		"john",
		&Address{"13 London", "London", "Uk"},
		[]string{"chaitanya1", "chaitanya 2"},
	}
	fmt.Println(john)

}

package main

import "fmt"

type Person struct {
	name, position string
}

type personMod func(person *Person)

type PersionBuilder struct {
	actions []personMod
}

func (b *PersionBuilder) Called(name string) *PersionBuilder {
	b.actions = append(b.actions, func(person *Person) {
		person.name = name
	})
	return b
}

func (b *PersionBuilder) Build() *Person {
	p := &Person{}
	for _, a := range b.actions {
		a(p)
	}
	return p
}

func (b *PersionBuilder) WorksAsa(position string) *PersionBuilder {
	b.actions = append(b.actions, func(person *Person) {
		person.position = position
	})
	return b
}

func main() {

	builder := PersionBuilder{}
	p := builder.Called("chaitanya").WorksAsa("Developer").Build()
	fmt.Println(p)
}

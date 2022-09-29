package main

import "fmt"

type Person struct {
	FirstName, MiddleName, LastnName string
}

func (p *Person) Names() [3]string {
	return [3]string{
		p.FirstName, p.MiddleName, p.LastnName,
	}
}

func (p *Person) NameGenerator() <-chan string {
	out := make(chan string)
	go func() {

		defer close(out)
		out <- p.FirstName
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastnName
	}()

	return out
}

type PersonIterator struct {
	person  *Person
	current int
}

func NewPersonIterator(person *Person) *PersonIterator {
	return &PersonIterator{person: person, current: -1}
}

func (p *PersonIterator) MoveNext() bool {
	p.current++
	return p.current < 3
}

func (p *PersonIterator) Value() string {
	switch p.current {
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2:
		return p.person.LastnName
	default:
		panic("we should not be in here")
	}
}

func main() {
	p := Person{"Alaxender", "Grahambell", "Bell"}
	for name := range p.NameGenerator() {
		fmt.Println(name)
	}
	for it := NewPersonIterator(&p); it.MoveNext(); {
		fmt.Println(it.Value())
	}

}

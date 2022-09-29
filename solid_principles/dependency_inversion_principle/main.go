package main

import "fmt"

// HLM should not depend on LLM
// Both should depend on abstractions

type RelationShip int

const (
	Parent RelationShip = iota
	Child
	Sibling
)

type Person struct {
	name string
}
type Info struct {
	from         *Person
	relationship RelationShip
	to           *Person
}

type Relationships struct {
	relations []Info
}

func (r *Relationships) FindAllChildren(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}

	return result
}

func (r *Relationships) AddParentoChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// high level module
type Research struct {

	// Break DIP
	relationships Relationships
}

func (r Research) Investigate() {
	relations := r.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "John" && rel.relationship == Parent {
			fmt.Println("John has a child", rel.to.name)
		}
	}
}

func main() {

	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"raj"}

	relationships := Relationships{}
	relationships.AddParentoChild(&parent, &child1)
	relationships.AddParentoChild(&parent, &child2)

	r := Research{relationships: relationships}
	r.Investigate()

	r2 := Research2{&relationships}
	r2.Investigate()
}

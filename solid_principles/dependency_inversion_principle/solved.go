package main

import "fmt"

type RelationshipBrowser interface {
	FindAllChildren(name string) []*Person
}

type Research2 struct {
	browser RelationshipBrowser
}

func (r *Research2) Investigate() {
	for _, p := range r.browser.FindAllChildren("John") {
		fmt.Println(p)
	}
}

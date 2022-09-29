package main

import "fmt"

// OCP
// Open for extension , closed for modification
// Specification pattern
type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f *Filter) filterByColor(products []Product, color Color) []*Product {

	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			fmt.Println(v)
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) filterBySize(products []Product, size Size) []*Product {

	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			fmt.Println(v)
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {

	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Println("Get all the Green Products old")

	f := Filter{}

	for _, v := range f.filterByColor(products, green) {
		fmt.Println(v.name)
	}

	greenSpec := ColorSpecification{green}
	f2 := BetterFilter{}
	for _, v := range f2.Filter(products, greenSpec) {

		fmt.Println(v)

	}
	fmt.Println("done")

	largeSpec := SizeSpecification{large}
	largegreenspec := AndSpecification{greenSpec, largeSpec}
	for _, v := range f2.Filter(products, largegreenspec) {
		fmt.Println(v)
	}

}

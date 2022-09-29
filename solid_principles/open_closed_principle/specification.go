package main

type Specification interface {
	isSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) isSatisfied(p *Product) bool {
	return p.color == c.color

}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) isSatisfied(p *Product) bool {
	return p.size == s.size

}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) isSatisfied(p *Product) bool {
	return a.first.isSatisfied(p) && a.second.isSatisfied(p)

}

type BetterFilter struct {
}

func (b *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.isSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

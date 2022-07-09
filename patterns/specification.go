package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
	brown
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

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type AndSpecification struct {
	left, right Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.left.IsSatisfied(p) && a.right.IsSatisfied(p)
}

type Filter struct{}

func (f *Filter) Filter(products []Product, spec Specification) []*Product {
	result := []*Product{}
	for i, p := range products {
		if spec.IsSatisfied(&p) {
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
	fmt.Printf("Products: %v\n", products)

	fmt.Println("Filtered:")
	f1 := Filter{}
	filtered := f1.Filter(products, AndSpecification{ColorSpecification{green}, SizeSpecification{large}})
	for _, v := range filtered {
		fmt.Printf("- %v\n", *v)
	}
}

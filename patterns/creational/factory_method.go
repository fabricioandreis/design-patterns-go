package creational

import "fmt"

// Factory Method (or Virtual Constructor or Factory Function) is a creational design pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created.
// Factory handles the wholesale creation of objects (unlike piecewise like a Builder)
// Objects returned by a Factory are called Products
// https://refactoring.guru/design-patterns/factory-method

type CustomPerson struct {
	Name     string
	Age      int
	EyeCount int
}

func NewCustomPerson(name string, age int) *CustomPerson {
	return &CustomPerson{name, age, 2}
}

type person struct {
	name string
	age  int
}

type PersonInterface interface {
	SayHello()
}

func (p *tiredPerson) SayHello() {
	fmt.Println("Sorry, I'm too tired")
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, I am %d years old\n", p.name, p.age)
}

func NewPerson(name string, age int) PersonInterface {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

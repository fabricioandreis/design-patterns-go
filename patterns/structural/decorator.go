package structural

import "fmt"

// Decorator is a structural design pattern that lets you attach new behaviors to objects by placing these objects inside special wrapper objects that contain the behaviors.
// https://refactoring.guru/design-patterns/decorator

// Want to augment an object with additional functionality
// Do not want to rewrite or alter existing code (OCP)
// Want to keep new functionality separate (SRP)
// Need to be able to interact with existing structures

// Solution: embed the decorated object and provide additional functionality

// The implementation bellow is not good because the composed DragonBroken struct
// has two Age fields within its embedded structs
type BirdBroken struct {
	Age int
}

func (b *BirdBroken) Fly() bool {
	if b.Age >= 10 {
		fmt.Println("Flying!")
		return true
	}
	return false
}

type LizardBroken struct {
	Age int
}

func (l *LizardBroken) Crawl() bool {
	if l.Age < 10 {
		fmt.Printf("Crawling!")
		return true
	}
	return false
}

type DragonBroken struct {
	BirdBroken
	LizardBroken
}

// This implementation uses multiple aggregation to properly implement a Dragon
type Aged interface {
	Age() int
	SetAge(age int) // Not idiomatic in Go, but we cannot avoid in this situation
}

type Bird struct {
	age int
}

func (b *Bird) Age() int {
	return b.age
}

func (b *Bird) SetAge(age int) {
	b.age = age
}

func (b *Bird) Fly() bool {
	if b.age >= 10 {
		fmt.Println("Flying!")
		return true
	}
	return false
}

type Lizard struct {
	age int
}

func (b *Lizard) Age() int {
	return b.age
}

func (b *Lizard) SetAge(age int) {
	b.age = age
}

func (l *Lizard) Crawl() bool {
	if l.age < 10 {
		fmt.Printf("Crawling!")
		return true
	}
	return false
}

type Dragon struct {
	bird   Bird
	lizard Lizard
}

func (b *Dragon) Age() int {
	return b.bird.age
}

func (b *Dragon) SetAge(age int) {
	b.bird.age = age
	b.lizard.age = age
}

func (d *Dragon) Fly() bool {
	return d.bird.Fly()
}

func (d *Dragon) Crawl() bool {
	return d.lizard.Crawl()
}

func NewDragon() *Dragon {
	return &Dragon{Bird{}, Lizard{}}
}

type ShapeDecor interface {
	Render() string
}

type CircleDecor struct {
	Radius float32
}

func (c *CircleDecor) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

func (c *CircleDecor) Resize(factor float32) {
	c.Radius *= factor
}

type SquareDecor struct {
	Side float32
}

func (c *SquareDecor) Render() string {
	return fmt.Sprintf("Square with side %f", c.Side)
}

type ColoredShapeDecor struct {
	Shape ShapeDecor
	Color string
}

func (c *ColoredShapeDecor) Render() string {
	return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}

type TransparentShapeDecor struct {
	Shape        ShapeDecor
	Transparency float32
}

func (t *TransparentShapeDecor) Render() string {
	return fmt.Sprintf("%s has %f%% transparency", t.Shape.Render(), t.Transparency*100.0)
}

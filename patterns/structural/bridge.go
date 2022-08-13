package structural

import "fmt"

// Bridge is a structural design pattern that lets you split a large class or a set
// of closely related classes into two separate hierarchies—abstraction and implementation—which
// can be developed independently of each other.
// https://refactoring.guru/design-patterns/bridge

// A mechanism that decouples an interface (hierarchy) from an implementation (hierarchy)

// Renderer is the Implementation interface
type Renderer interface {
	RenderCircle(radius float64)
}

// VectorRenderer is a Concrete Implementation
type VectorRenderer struct {
}

func (v *VectorRenderer) RenderCircle(radius float64) {
	fmt.Println("Drawing a circle of radius", radius)
}

// RasterRenderer is a Concrete Implementation
type RasterRenderer struct {
	Dpi int
}

func (r *RasterRenderer) RenderCircle(radius float64) {
	fmt.Println("Drawing pixels for a circle of radius", radius)
}

// Circle is the Abstraction
type Circle struct {
	renderer Renderer
	radius   float64
}

func NewCircle(renderer Renderer, radius float64) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float64) {
	c.radius *= factor
}

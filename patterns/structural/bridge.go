package structural

// Bridge is a structural design pattern that lets you split a large class or a set
// of closely related classes into two separate hierarchies—abstraction and implementation—which
// can be developed independently of each other.
// https://refactoring.guru/design-patterns/bridge

// A mechanism that decouples an interface (hierarchy) from an implementation (hierarchy)

type Renderer interface {
	RenderCircle(radius float64)
}

type VectorRenderer struct {
}

func (v *VectorRenderer) RenderCircle(radius float64) {

}

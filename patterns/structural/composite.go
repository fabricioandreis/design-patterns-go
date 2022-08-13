package structural

import (
	"strings"
)

// Composite is a structural design pattern that lets you compose objects into tree structures and then work with these structures as if they were individual objects.
// A mechanism for treating individual (scalar) objects and compositions of objects in a uniform manner.
// https://refactoring.guru/design-patterns/composite

type GraphicObject struct {
	Name, Color string
	Children    []GraphicObject
}

func (g *GraphicObject) String() string {
	sb := strings.Builder{}
	g.print(&sb, 0)
	return sb.String()
}

func (g *GraphicObject) print(sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat("*", depth))
	if len(g.Color) > 0 {
		sb.WriteString(g.Color)
		sb.WriteRune(' ')
	}
	sb.WriteString(g.Name)
	sb.WriteRune('\n')
	for _, child := range g.Children {
		child.print(sb, depth+1)
	}
}

func NewCompositeCircle(color string) *GraphicObject {
	return &GraphicObject{"Circle", color, nil}
}

func NewCompositeSquare(color string) *GraphicObject {
	return &GraphicObject{"Square", color, nil}
}

type Component interface {
	All() []*Neuron
}

// We want to connect neurons to other neurons, neuron to neuron layers, and neuron layers to neuron layers
type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) All() []*Neuron {
	return []*Neuron{n}
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}

type NeuronLayer struct {
	Neurons []Neuron
}

func (l *NeuronLayer) All() []*Neuron {
	result := []*Neuron{}
	for i := range l.Neurons {
		result = append(result, &l.Neurons[i])
	}
	return result
}

func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{make([]Neuron, count)}
}

func Connect(left, right Component) {
	for _, l := range left.All() {
		for _, r := range right.All() {
			l.ConnectTo(r)
		}
	}
}

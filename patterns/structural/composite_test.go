package structural_test

import (
	"fmt"
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/structural"
)

func TestComposite(t *testing.T) {
	t.Run("Should print a hierarchy of objects", func(t *testing.T) {
		g1 := structural.GraphicObject{Name: "Circle", Color: "Blue"}
		g2 := structural.GraphicObject{Name: "Square", Color: "Red", Children: []structural.GraphicObject{g1}}

		fmt.Println(g2.String())
	})

	t.Run("Should print a complex hierarchy of objects", func(t *testing.T) {
		drawing := structural.GraphicObject{Name: "Drawing", Color: "", Children: nil}

		drawing.Children = append(drawing.Children, *structural.NewCompositeCircle("Red"))
		drawing.Children = append(drawing.Children, *structural.NewCompositeSquare("Yellow"))

		group1 := structural.GraphicObject{Name: "Group 1", Color: "", Children: nil}
		group1.Children = append(group1.Children, *structural.NewCompositeCircle("Blue"))
		group1.Children = append(group1.Children, *structural.NewCompositeSquare("Blue"))
		drawing.Children = append(drawing.Children, group1)

		fmt.Println(drawing.String())
	})

	t.Run("Should create Neuron Layer", func(t *testing.T) {
		structural.NewNeuronLayer(10)
	})

	t.Run("Should connect Neurons and Neuron Layers", func(t *testing.T) {
		neuron1, neuron2 := &structural.Neuron{}, &structural.Neuron{}
		layer1, layer2 := structural.NewNeuronLayer(3), structural.NewNeuronLayer(4)

		structural.Connect(neuron1, neuron2)
		structural.Connect(neuron1, layer1)
		structural.Connect(layer2, neuron1)
		structural.Connect(layer1, layer2)
	})
}

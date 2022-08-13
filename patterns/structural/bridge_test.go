package structural_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/structural"
)

func TestBridge(t *testing.T) {

	t.Run("Should draw a circle with vector renderer", func(t *testing.T) {
		v := structural.VectorRenderer{}
		c := structural.NewCircle(&v, 100)

		c.Draw()
	})

	t.Run("Should draw a circle with raster renderer", func(t *testing.T) {
		v := structural.RasterRenderer{}
		c := structural.NewCircle(&v, 100)

		c.Draw()
	})

	t.Run("Should draw a resized circle with raster renderer", func(t *testing.T) {
		v := structural.RasterRenderer{}
		c := structural.NewCircle(&v, 100)

		c.Draw()
		c.Resize(200)
		c.Draw()
	})

}

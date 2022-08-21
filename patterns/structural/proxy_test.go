package structural_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/structural"
	"github.com/stretchr/testify/assert"
)

func TestProxy(t *testing.T) {
	t.Run("Underage driver should not drive SafeCar", func(t *testing.T) {
		tests := map[int]bool{
			17: false,
			14: false,
			18: true,
			34: true,
		}

		for age, output := range tests {
			driver := structural.Driver{age}
			car := structural.NewSafeCar(&driver)

			driving := car.Drive()

			assert.Equal(t, output, driving)
		}
	})

	t.Run("Should draw an image with bitmap", func(t *testing.T) {
		bmp := structural.NewBitmap("demo.png")
		structural.DrawImage(bmp)
		// Should print:
		// Loading image from demo.png
		// About to draw the image
		// Drawing image demo.png
		// Done drawing the image
	})

	t.Run("Should draw an image with lazy bitmap", func(t *testing.T) {
		bmp := structural.NewLazyBitmap("demo.png")
		structural.DrawImage(bmp)
		// Should print:
		// About to draw the image
		// Loading image from demo.png
		// Drawing image demo.png
		// Done drawing the image
	})
}

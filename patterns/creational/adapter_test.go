package creational_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/creational"
	"github.com/stretchr/testify/assert"
)

func TestAdapter(t *testing.T) {
	t.Run("Should create rectangle vector image", func(t *testing.T) {
		width := 100
		height := 150
		r := creational.NewRectangle(width, height)

		assert.NotNil(t, r)
		assert.Equal(t, 0, r.Lines[0].X1)
		assert.Equal(t, 0, r.Lines[0].Y1)
		assert.Equal(t, width-1, r.Lines[0].X2)
		assert.Equal(t, 0, r.Lines[0].Y2)
	})

	t.Run("Should draw default (raster) image", func(t *testing.T) {
		p1 := creational.Point{0, 2}
		p2 := creational.Point{2, 3}
		img := creational.DefaultImage{[]creational.Point{p1, p2}}

		output := creational.DrawPoints(&img)

		assert.NotEqual(t, "", output)
		assert.Equal(t, "   \n   \n*  \n  *\n", output)
	})

	t.Run("Should convert rectangle vector image to raster image", func(t *testing.T) {
		r := creational.NewRectangle(5, 3)

		img := creational.VectorToRaster(r)
		output := creational.DrawPoints(img)

		assert.NotEqual(t, "", output)
		assert.Equal(t, "*****\n*   *\n*****\n", output)
	})

	t.Run("Should cover cache usage", func(t *testing.T) {
		r := creational.NewRectangle(5, 3)

		img := creational.VectorToRaster(r)
		creational.VectorToRaster(r)
		output := creational.DrawPoints(img)

		assert.NotEqual(t, "", output)
		assert.Equal(t, "*****\n*   *\n*****\n", output)
	})
}

package principles_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/principles"
	"github.com/stretchr/testify/assert"
)

func TestLSP(t *testing.T) {

	t.Run("Should work for a rectangle", func(t *testing.T) {
		rectangle := principles.Rectangle{2, 3}

		width := rectangle.Width()
		rectangle.SetHeight(10)

		assert.Equal(t, 10*width, rectangle.Width()*rectangle.Height())
	})

	t.Run("Should not work for a square", func(t *testing.T) {
		square := principles.NewSquare(5)

		width := square.Width()
		square.SetHeight(10)

		assert.NotEqual(t, 10*width, square.Width()*square.Height())
	})

}

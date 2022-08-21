package structural_test

import (
	"fmt"
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/structural"
	"github.com/stretchr/testify/assert"
)

func TestMultipleAggregation(t *testing.T) {
	t.Run("Bird should not fly when less than age 10", func(t *testing.T) {
		tests := map[int]bool{
			8:  false,
			7:  false,
			10: true,
			11: true,
		}

		for age, output := range tests {
			t.Run(fmt.Sprintf("Testing age %v", age), func(t *testing.T) {
				b := structural.BirdBroken{Age: age}

				fly := b.Fly()

				assert.Equal(t, output, fly)
			})
		}
	})

	t.Run("Lizard should crawl when less than age 10", func(t *testing.T) {
		tests := map[int]bool{
			8:  true,
			7:  true,
			10: false,
			11: false,
		}

		for age, output := range tests {
			t.Run(fmt.Sprintf("Testing age %v", age), func(t *testing.T) {
				b := structural.LizardBroken{Age: age}

				crawl := b.Crawl()

				assert.Equal(t, output, crawl)
			})
		}
	})

	t.Run("Need to set ages separately", func(t *testing.T) {
		d := structural.DragonBroken{}

		d.BirdBroken.Age = 10
		d.LizardBroken.Age = 10

		fly := d.Fly()

		assert.Equal(t, true, fly)
	})

	t.Run("Works fine for correct Dragon", func(t *testing.T) {
		d := structural.NewDragon()
		d.SetAge(10)

		fly := d.Fly()
		crawl := d.Crawl()

		assert.Equal(t, true, fly)
		assert.Equal(t, false, crawl)
	})
}

func TestDecorator(t *testing.T) {
	t.Run("Should decorate shape with color", func(t *testing.T) {
		circle := structural.CircleDecor{2}
		redCircle := structural.ColoredShapeDecor{&circle, "Red"}

		r := redCircle.Render()

		assert.Equal(t, "Circle of radius 2.000000 has the color Red", r)
	})
	t.Run("Should decorate a decorator", func(t *testing.T) {
		circle := structural.CircleDecor{2}
		redCircle := structural.ColoredShapeDecor{&circle, "Red"}
		rhsCircle := structural.TransparentShapeDecor{&redCircle, 0.5}

		r := rhsCircle.Render()

		assert.Equal(t, "Circle of radius 2.000000 has the color Red has 50.000000% transparency", r)
	})
}

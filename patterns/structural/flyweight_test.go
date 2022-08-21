package structural_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/structural"
	"github.com/stretchr/testify/assert"
)

func TestFlyweight(t *testing.T) {
	t.Run("Should capitalize in heavy weight (memory inefficient)", func(t *testing.T) {
		text := "Yet though I walk through the valley of the shadow of death, I will fear no evil for thou art with me"

		ft := structural.NewFormattedTextHeavy(text)
		ft.Capitalize(62, 100)
		capitalized := ft.String()

		assert.Equal(
			t,
			"Yet though I walk through the valley of the shadow of death, I WILL FEAR NO EVIL FOR THOU ART WITH ME",
			capitalized)
	})

	t.Run("Should capitalize in flyweight (memory efficient)", func(t *testing.T) {
		text := "Yet though I walk through the valley of the shadow of death, I will fear no evil for thou art with me"

		ft := structural.NewFormattedText(text)
		ft.Range(62, 100).Capitalize = true
		capitalized := ft.String()

		assert.Equal(
			t,
			"Yet though I walk through the valley of the shadow of death, I WILL FEAR NO EVIL FOR THOU ART WITH ME",
			capitalized)
	})

	t.Run("Should create users in heavy weight (memory inefficient)", func(t *testing.T) {
		structural.NewUserHeavy("John Doe")
		structural.NewUserHeavy("Jane Doe")
		structural.NewUserHeavy("Jane Smith")
	})

	t.Run("Should create users in flyweight (memory efficient)", func(t *testing.T) {
		john := structural.NewUser("John Doe")
		jane1 := structural.NewUser("Jane Doe")
		jane2 := structural.NewUser("Jane Smith")

		assert.Equal(t, "John Doe", john.FullName())
		assert.Equal(t, "Jane Doe", jane1.FullName())
		assert.Equal(t, "Jane Smith", jane2.FullName())
	})
}

package creational_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/creational"
	"github.com/stretchr/testify/assert"
)

func TestFactoryMethod(t *testing.T) {
	t.Run("Should make a product with a Factory Method", func(t *testing.T) {
		p := creational.NewCustomPerson("Rafael", 1)

		assert.Equal(t, "Rafael", p.Name)
		assert.Equal(t, 1, p.Age)
		assert.Equal(t, 2, p.EyeCount)
	})

	t.Run("Should make a person and print message", func(t *testing.T) {
		p := creational.NewPerson("Rafael", 1)

		p.SayHello()
	})
	t.Run("Should make a tired person and print message", func(t *testing.T) {
		p := creational.NewPerson("John", 101)

		p.SayHello()
	})
}

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

	t.Run("Should generate custom factories and create products from them", func(t *testing.T) {
		developerFactory := creational.NewEmployeeFactory("developer", 60000)
		managerFactory := creational.NewEmployeeFactory("manager", 80000)

		d := developerFactory("Adam")
		m := managerFactory("Jane")

		assert.Equal(t, "Adam", d.Name)
		assert.Equal(t, "developer", d.Position)
		assert.Equal(t, 60000, d.AnnualIncome)
		assert.Equal(t, "Jane", m.Name)
		assert.Equal(t, "manager", m.Position)
		assert.Equal(t, 80000, m.AnnualIncome)
	})

	t.Run("Should instantiate a factory and create products from it", func(t *testing.T) {
		bossFactory := creational.NewEmployeeFactoryStruct("boss", 100000)

		b := bossFactory.Create("James")

		assert.Equal(t, "James", b.Name)
		assert.Equal(t, "boss", b.Position)
		assert.Equal(t, 100000, b.AnnualIncome)
	})

	t.Run("Should allow to change a factory and create products from it", func(t *testing.T) {
		bossFactory := creational.NewEmployeeFactoryStruct("boss", 100000)
		bossFactory.Position = "slave"
		bossFactory.AnnualIncome = 0
		b := bossFactory.Create("Moses")

		assert.Equal(t, "Moses", b.Name)
		assert.Equal(t, "slave", b.Position)
		assert.Equal(t, 0, b.AnnualIncome)
	})
}

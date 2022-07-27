package creational_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/creational"
	"github.com/stretchr/testify/assert"
)

func TestCreateParagraph(t *testing.T) {
	value := "hello"

	par := creational.CreateParagraph(value)

	assert.Equal(t, "<p>hello</p>", par)
}

func TestHtmlBuilder(t *testing.T) {
	t.Run("Should work for regular build step method", func(t *testing.T) {
		b := creational.NewHtmlBuilder("ul")
		b.AddChild("li", "hello")
		b.AddChild("li", "world")

		t.Log(b.String())
	})

	t.Run("Should also work for fluent build step method", func(t *testing.T) {
		b := creational.NewHtmlBuilder("ul").AddChildFluent("li", "hello").AddChildFluent("li", "world")
		t.Log(b.String())
	})
}

func TestBuilderFacets(t *testing.T) {
	t.Run("Should build an object with builder facets", func(t *testing.T) {
		builder := creational.NewPersonBuilder()
		builder.
			Lives().
			At("Main street").
			In("Caxias do Sul").
			WithPostCode("95000000").
			Works().
			At("Company").
			AsA("Engineer").
			Earning(100000)

		p := builder.Build()

		assert.Equal(t, "Main street", p.StreetAddress)
		assert.Equal(t, "Caxias do Sul", p.City)
		assert.Equal(t, "95000000", p.PostCode)
		assert.Equal(t, "Company", p.CompanyName)
		assert.Equal(t, "Engineer", p.Position)
		assert.Equal(t, 100000, p.AnnualIncome)
	})
}

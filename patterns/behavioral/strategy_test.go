package behavioral_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/behavioral"
	"github.com/stretchr/testify/assert"
)

// Strategy is a behavioral design pattern that lets you define a family of algorithms, put each of them into a separate class, and make their objects interchangeable.

// https://refactoring.guru/design-patterns/strategy

// Separates an algorithm into its 'skeleton' and concrete implementation steps, which can be varied at runtime

// Many algorithms can be decomposed into higher and lower level parts.
// E.g. making a tea can be decomposed into
// 1. The process of making a hot beverage (boil water, pour into cup); and
// 2. Tea-specific things (put teabag into water)
// The high-level algorithm can then be reused for making coffee of hot chocolate

func TestStrategy(t *testing.T) {
	t.Run("Should be able to switch strategies at runtime", func(t *testing.T) {
		tp := behavioral.NewTextProcessor(behavioral.Html)

		input := [][]string{
			{"item 1", "item 2"},
			{"item 1", "item 2"},
			{"item 3", "item 2", "item 1"},
			{"item 3", "item 2", "item 1"},
		}
		strategy := []behavioral.OutputFormat{
			behavioral.Markdown,
			behavioral.Html,
			behavioral.Html,
			behavioral.Markdown,
		}
		output := []string{
			" * item 1\n * item 2\n",
			"<ul>\n  <li>item 1</li>\n  <li>item 2</li>\n</ul>\n",
			"<ul>\n  <li>item 3</li>\n  <li>item 2</li>\n  <li>item 1</li>\n</ul>\n",
			" * item 3\n * item 2\n * item 1\n",
		}

		for i, input := range input {
			tp.Reset()
			tp.SetOutputFormat(strategy[i])
			tp.AppendList(input)
			result := tp.String()

			assert.Equal(t, output[i], result)
		}
	})
}

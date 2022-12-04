package behavioral_test

import (
	"strings"
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/behavioral"
	"github.com/stretchr/testify/assert"
)

func TestVisitor(t *testing.T) {
	t.Run("Should visit elements intrusively", func(t *testing.T) {
		// Evaluating 1 + (2 + 3)
		e := &behavioral.AdditionExpression{
			Left: &behavioral.DoubleExpression{1},
			Right: &behavioral.AdditionExpression{
				Left:  &behavioral.DoubleExpression{2},
				Right: &behavioral.DoubleExpression{3},
			},
		}
		sb := strings.Builder{} // Visitor

		e.Print(&sb)
		output := sb.String()

		assert.Equal(t, "(1+(2+3))", output)
	})

	t.Run("Should visit elements using visitor that checks types", func(t *testing.T) {
		// Evaluating 1 + (2 + 3)
		e := &behavioral.AdditionExpression{
			Left: &behavioral.DoubleExpression{1},
			Right: &behavioral.AdditionExpression{
				Left:  &behavioral.DoubleExpression{2},
				Right: &behavioral.DoubleExpression{3},
			},
		}
		sb := strings.Builder{} // Visitor

		behavioral.PrintExpression(e, &sb)
		output := sb.String()

		assert.Equal(t, "(1+(2+3))", output)
	})

	t.Run("Should visit elements using visitor with double dispatch", func(t *testing.T) {
		// Evaluating 1 + (2 + 3)
		e := &behavioral.AdditionExpression{
			Left: &behavioral.DoubleExpression{1},
			Right: &behavioral.AdditionExpression{
				Left:  &behavioral.DoubleExpression{2},
				Right: &behavioral.DoubleExpression{3},
			},
		}
		ep := behavioral.NewExpressionPrinter() // Visitor

		e.Accept(ep)
		output := ep.String()

		assert.Equal(t, "(1+(2+3))", output)
	})

	t.Run("Should visit elements using a new visitor with double dispatch", func(t *testing.T) {
		// Evaluating 1 + (2 + 3)
		e := &behavioral.AdditionExpression{
			Left: &behavioral.DoubleExpression{1},
			Right: &behavioral.AdditionExpression{
				Left:  &behavioral.DoubleExpression{2},
				Right: &behavioral.DoubleExpression{3},
			},
		}
		ee := behavioral.NewExpressionEvaluator() // Visitor

		e.Accept(ee)
		output := ee.Result()

		assert.Equal(t, 6.0, output)
	})
}

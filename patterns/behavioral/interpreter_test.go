package behavioral_test

import (
	"fmt"
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/behavioral"
)

func TestInterpreter(t *testing.T) {
	t.Run("Should be able to perform lexing of a numeric expression", func(t *testing.T) {
		input := "(13+4)-(12+1)"
		tokens := behavioral.Lex(input)
		fmt.Println(tokens)
	})

	t.Run("Should be able to perform parsing of a numeric expression", func(t *testing.T) {
		input := "(13+4)-(12+1)"
		tokens := behavioral.Lex(input)
		element := behavioral.Parse(tokens)
		fmt.Printf("%s = %v", input, element.Value())
	})
}

package structural_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/structural"
)

func TestFacade(t *testing.T) {
	t.Run("Should use simpler Console Facade instead of buffers and viewports", func(t *testing.T) {
		c := structural.NewConsole()

		c.GetCharacterAt(10)
	})
}

package behavioral_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/behavioral"
)

func TestState(t *testing.T) {
	t.Run("Should be able to transition a state", func(t *testing.T) {
		sw := behavioral.NewSwitch()
		sw.On()
		sw.Off()
		sw.Off()
	})

	t.Run("Should be able to transition a states with simpler structures", func(t *testing.T) {
		sw := behavioral.NewSwitch()
		sw.On()
		sw.Off()
		sw.Off()
	})
}

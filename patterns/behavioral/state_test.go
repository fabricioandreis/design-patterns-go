package behavioral_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/behavioral"
	"github.com/stretchr/testify/assert"
)

func TestState(t *testing.T) {
	t.Run("Should be able to transition a state", func(t *testing.T) {
		sw := behavioral.NewSwitch()
		sw.On()
		sw.Off()
		sw.Off()
	})

	t.Run("Should be able to transition states from a landline phone", func(t *testing.T) {
		state := behavioral.OffHook
		tr := behavioral.Rules[state][0] // CallDialed: Connecting
		state = tr.PhoneState
		tr = behavioral.Rules[state][1] // CallConnected: Connected
		state = tr.PhoneState
		tr = behavioral.Rules[state][0] // LeftMessage: OnHook
		state = tr.PhoneState

		assert.Equal(t, behavioral.LeftMessage, tr.PhoneTrigger)
		assert.Equal(t, behavioral.OnHook, state)
	})

	t.Run("Should be able to transition states with switch statement", func(t *testing.T) {
		type testResult struct {
			output bool
			state  behavioral.SystemState
		}
		sys := behavioral.NewSystemSecret("Rafael")
		tests := map[string]testResult{
			"Daiana":   {false, behavioral.Failed},
			"Fabrício": {false, behavioral.Locked},
			"Rafael":   {true, behavioral.Unlocked},
		}

		for code, result := range tests {
			output := sys.Unlock(code)

			assert.Equal(t, result.output, output)
			assert.Equal(t, result.state, sys.State())
		}
	})
}

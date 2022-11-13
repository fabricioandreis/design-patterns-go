package behavioral_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/behavioral"
	"github.com/stretchr/testify/assert"
)

func TestMemento(t *testing.T) {
	t.Run("Return bank account to previous state", func(t *testing.T) {
		ba, mem0 := behavioral.NewBankAccountMemento(50.0)
		mem1 := ba.Deposit(100.0)
		ba.Deposit(200.0)
		ba.Restore(mem1)
		assert.Equal(t, 150.0, ba.Balance())
		ba.Restore(mem0)
		assert.Equal(t, 50.0, ba.Balance())
	})

	t.Run("Should be able to do and redo actions in payment account", func(t *testing.T) {
		pa := behavioral.NewPaymentAccount(25.0)
		pa.Deposit(100.0)
		pa.Deposit(200.0)

		pa.Undo()
		assert.Equal(t, 125.0, pa.Balance())
		pa.Undo()
		assert.Equal(t, 25.0, pa.Balance())
		pa.Redo()
		assert.Equal(t, 125.0, pa.Balance())
	})
}

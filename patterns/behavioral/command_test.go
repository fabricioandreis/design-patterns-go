package behavioral_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/behavioral"
	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	t.Run("Should execute commands", func(t *testing.T) {
		ba := behavioral.BankAccount{}
		cmd1 := behavioral.NewBankAccountCommand(&ba, behavioral.Deposit, 100)
		cmd2 := behavioral.NewBankAccountCommand(&ba, behavioral.Withdraw, 50)

		cmd1.Call()
		cmd2.Call()

		assert.Equal(t, ba.Balance(), 50)
	})

	t.Run("Should be able to undo commands", func(t *testing.T) {

	})
}

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
		ba := behavioral.BankAccount{}
		cmd1 := behavioral.NewBankAccountCommand(&ba, behavioral.Deposit, 100)
		cmd2 := behavioral.NewBankAccountCommand(&ba, behavioral.Withdraw, 25)

		cmd1.Call()
		cmd2.Call()
		cmd2.Undo()
		cmd1.Undo()

		assert.Equal(t, ba.Balance(), 0)
	})

	t.Run("Should be able to transfer money atomically", func(t *testing.T) {
		from := &behavioral.BankAccount{}
		from.Deposit(100)
		to := &behavioral.BankAccount{}
		to.Deposit(0)

		mtc := behavioral.NewMoneyTransferCommand(from, to, 25)
		mtc.Call()

		assert.Equal(t, from.Balance(), 75)
		assert.Equal(t, to.Balance(), 25)
	})

	t.Run("Should be able to undo atomic money transfer", func(t *testing.T) {
		from := &behavioral.BankAccount{}
		from.Deposit(100)
		to := &behavioral.BankAccount{}
		to.Deposit(0)

		mtc := behavioral.NewMoneyTransferCommand(from, to, 25)
		mtc.Call()
		mtc.Undo()

		assert.Equal(t, from.Balance(), 100)
		assert.Equal(t, to.Balance(), 0)
	})

	t.Run("Should be able to perform commands in a functional approach", func(t *testing.T) {
		ba := behavioral.BankAccount{}
		var commands []func()
		commands = append(commands, func() {
			ba.Deposit(100)
		})
		commands = append(commands, func() {
			ba.Withdraw(25)
		})
		for _, cmd := range commands {
			cmd()
		}

		assert.Equal(t, ba.Balance(), 75)
	})
}

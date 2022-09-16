package behavioral

import (
	"fmt"
)

// Command is a behavioral design pattern that turns a request into a stand-alone object that contains all information about the request.
// This transformation lets you pass requests as a method arguments, delay or queue a request’s execution, and support undoable operations.
// A command object represents an instruction to perform a particular action. Contains all the information necessary for the action to be taken.
// Uses: GUI commands, multi-level undo/redo, macro recording and more.
// Commands can be processed by the object upon which it performs its action, by itself or by another processor object (service)

// https://refactoring.guru/design-patterns/command

const overdraftLimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) bool {
	b.balance += amount
	fmt.Println("Deposited", amount, "\b, balance is now", b.balance)
	return true
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		fmt.Println("Withdrew", amount, "\b, balance is now", b.balance)
		return true
	}
	return false
}

func (b *BankAccount) Balance() int {
	return b.balance
}

type Command interface {
	Call()
	Undo()
	Succeeded() bool
	SetSucceeded(value bool)
}

type action int

const (
	Deposit action = iota
	Withdraw
)

type Action interface {
	Action() action
}

func (a action) Action() action {
	return a
}

type BankAccountCommand struct {
	account   *BankAccount
	action    Action
	amount    int
	succeeded bool
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account, action, amount, false}
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.succeeded = b.account.Deposit(b.amount)
	case Withdraw:
		b.succeeded = b.account.Withdraw(b.amount)
	}
}

func (b *BankAccountCommand) Undo() {
	if !b.succeeded {
		return
	}
	switch b.action {
	case Deposit:
		b.account.Withdraw(b.amount)
	case Withdraw:
		b.account.Deposit(b.amount)
	}
}

func (b *BankAccountCommand) Succeeded() bool {
	return b.succeeded
}

func (b *BankAccountCommand) SetSucceeded(value bool) {
	b.succeeded = value
}

type CompositeBankAccountCommand struct {
	commands []Command
}

func (c *CompositeBankAccountCommand) Call() {
	for _, c := range c.commands {
		c.Call()
	}
}

func (c *CompositeBankAccountCommand) Undo() {
	for i := range c.commands {
		c.commands[len(c.commands)-i-1].Undo()
	}
}

func (c *CompositeBankAccountCommand) Succeeded() bool {
	for _, c := range c.commands {
		if !c.Succeeded() {
			return false
		}
	}
	return true
}

func (c *CompositeBankAccountCommand) SetSucceeded(value bool) {
	for _, c := range c.commands {
		c.SetSucceeded(value)
	}
}

type MoneyTransferCommand struct {
	CompositeBankAccountCommand
	from, to *BankAccount
	amount   int
}

func NewMoneyTransferCommand(from *BankAccount, to *BankAccount, amount int) *MoneyTransferCommand {
	c := &MoneyTransferCommand{from: from, to: to, amount: amount}
	c.commands = append(c.commands, NewBankAccountCommand(from, Withdraw, amount))
	c.commands = append(c.commands, NewBankAccountCommand(to, Deposit, amount))
	return c
}

func (m *MoneyTransferCommand) Call() {
	ok := true
	for _, c := range m.commands {
		if ok {
			c.Call()
			ok = c.Succeeded()
		} else {
			c.SetSucceeded(false)
		}
	}
}

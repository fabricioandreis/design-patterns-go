package behavioral

// Memento is a behavioral design pattern that lets you save and restore the previous state of an object without revealing the details of its implementation.
// It is implemented as a token representing the system state. Lets us roll back to the state when the token was generated.
// May or may not directly expose state information

// https://refactoring.guru/design-patterns/memento

// Summary:
// - Mementos are used to roll back states arbitrarily
// - A memento is simply a token/handle with (typically) no methods of its own
// - A memento is not required to expose directly the states(s) to which it reverts the system
// - Can be used to implement undo/redo (but the Command design pattern saves memory)

type BankAccountMemento struct {
	balance float64
}

func NewBankAccountMemento(balance float64) (*BankAccountMemento, *Memento) {
	return &BankAccountMemento{balance}, &Memento{balance}
}

type Memento struct {
	Balance float64
}

func (b *BankAccountMemento) Deposit(amount float64) *Memento {
	b.balance += amount
	return &Memento{float64(b.balance)}
}

func (b *BankAccountMemento) Restore(m *Memento) {
	b.balance = m.Balance
}

func (b *BankAccountMemento) Balance() float64 {
	return b.balance
}

type PaymentAccount struct {
	balance float64
	changes []*Memento
	current int // index of the changes slice
}

func NewPaymentAccount(balance float64) *PaymentAccount {
	mem := &Memento{balance}
	return &PaymentAccount{
		balance: balance,
		changes: []*Memento{mem},
		current: 0,
	}
}

func (acc *PaymentAccount) Balance() float64 {
	return acc.balance
}

func (acc *PaymentAccount) Deposit(amount float64) *Memento {
	mem := &Memento{Balance: acc.balance + amount}
	acc.applyMemento(mem)
	return mem
}

func (acc *PaymentAccount) Restore(mem *Memento) {
	if mem == nil {
		return
	}

	acc.applyMemento(mem)
}

func (acc *PaymentAccount) applyMemento(mem *Memento) {
	acc.balance = mem.Balance
	acc.changes = append(acc.changes, mem)
	acc.current++
}

func (acc *PaymentAccount) Undo() *Memento {
	if acc.current <= 0 {
		return nil
	}
	acc.current--
	mem := acc.changes[acc.current]
	acc.balance = mem.Balance
	return mem
}

func (acc *PaymentAccount) Redo() *Memento {
	if acc.current+1 >= len(acc.changes) {
		return nil
	}
	acc.current++
	mem := acc.changes[acc.current]
	acc.balance = mem.Balance
	return mem
}

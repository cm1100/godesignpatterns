package main

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposited Amount ", amount, " balance is ", b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		fmt.Println("Withdrew Amount ", amount, " balance is ", b.balance)
		return true

	}
	return false
}

type Command interface {
	Call()
	Undo()
}

type Action int

const (
	Deposit Action = iota
	Widthraw
)

type BankAccountCommand struct {
	account  *BankAccount
	action   Action
	amount   int
	succeded bool
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
		b.succeded = true
	case Widthraw:
		b.succeded = b.account.Withdraw(b.amount)

	}
}

func (b *BankAccountCommand) Undo() {
	if !b.succeded {
		return
	}
	switch b.action {
	case Deposit:
		b.account.Withdraw(b.amount)
	case Widthraw:
		b.account.Deposit(b.amount)
	}
}

func main() {
	ba := &BankAccount{}
	cmd := NewBankAccountCommand(ba, Deposit, 100)
	cmd2 := NewBankAccountCommand(ba, Widthraw, 50)
	cmd.Call()
	cmd2.Call()
	cmd.Undo()
	cmd2.Undo()

	fmt.Println(ba)

}

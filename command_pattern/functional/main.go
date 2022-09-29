package main

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
	balance int
}

func Deposit(b *BankAccount, amount int) {
	b.balance += amount
	fmt.Println("Deposited Amount ", amount, " balance is ", b.balance)
}

func Withdraw(b *BankAccount, amount int) bool {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		fmt.Println("Withdrew Amount ", amount, " balance is ", b.balance)
		return true

	}
	return false
}

func main() {
	ba := &BankAccount{0}
	var commands []func()
	commands = append(commands, func() {
		Deposit(ba, 100)

	})
	commands = append(commands, func() {
		Withdraw(ba, 2)
	})

	for _, cmd := range commands {
		cmd()
	}
	fmt.Println(ba)
}

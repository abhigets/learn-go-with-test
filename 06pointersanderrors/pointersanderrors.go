package _6pointersanderrors

import (
	"errors"
	"fmt"
)

type BitCoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance BitCoin
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (wallet *Wallet) Deposit(amount BitCoin) {
	wallet.balance += amount
}

func (wallet *Wallet) Balance() interface{} {
	return wallet.balance
}

var InsufficientBalance = errors.New("insufficient balance")

func (wallet *Wallet) Withdraw(amount BitCoin) error {
	if amount > wallet.balance{
		return InsufficientBalance
	}

	wallet.balance -= amount
	return nil
}



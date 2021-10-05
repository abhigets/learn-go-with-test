package _6pointersanderrors

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("i am able to view my balance", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(BitCoin(10))
		actual:= wallet.Balance()
		expected := BitCoin(10)

		assert.Equal(t,actual,expected)
	})

	t.Run("i am able to withdraw from my balance", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(20)}

		wallet.Withdraw(BitCoin(10))
		actual:= wallet.Balance()
		expected := BitCoin(10)

		assert.Equal(t,actual,expected)
	})

	t.Run("i am not able to withdraw amount greater than my balance", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(30)}

		err := wallet.Withdraw(BitCoin(40))

		assert.Equal(t,err, InsufficientBalance)
		assert.Equal(t,InsufficientBalance.Error(), "insufficient balance")
		assert.Equal(t,wallet.balance,BitCoin(30))
	})
}

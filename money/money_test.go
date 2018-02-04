package money_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takkyuuplayer/tdd-book/money"
)

func TestMultipulation(t *testing.T) {
	five := money.Doller(5)

	assert.Equal(t, money.Doller(10), five.Times(2))
	assert.Equal(t, money.Doller(15), five.Times(3))
}

func TestEquality(t *testing.T) {
	assert.Equal(t, money.Doller(5).Equals(money.Doller(5)), true)
	assert.Equal(t, money.Doller(5).Equals(money.Doller(6)), false)

	assert.Equal(t, money.Doller(5).Equals(money.Franc(5)), false)
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, money.Doller(5).Currency(), "USD")
	assert.Equal(t, money.Franc(5).Currency(), "CHF")
}

func TestSimpleAddition(t *testing.T) {
	sum := money.Doller(5).Plus(money.Doller(5))

	bank := money.Bank{}
	reduced := bank.Reduce(sum, "USD")

	assert.Equal(t, reduced, money.Doller(10))
}

func TestPlusReturnsSum(t *testing.T) {
	five := money.Doller(5)
	result := five.Plus(five)

	sum := result.(*money.Sum)

	assert.Equal(t, five, sum.Augend)
	assert.Equal(t, five, sum.Addend)
}

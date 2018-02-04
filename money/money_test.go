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

	bank := money.NewBank()
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

func TestReduceSum(t *testing.T) {
	sum := &money.Sum{
		money.Doller(3),
		money.Doller(4),
	}
	bank := money.NewBank()
	result := bank.Reduce(sum, "USD")

	assert.Equal(t, money.Doller(7), result)
}

func TestReduceMoney(t *testing.T) {
	bank := money.NewBank()
	result := bank.Reduce(money.Doller(1), "USD")

	assert.Equal(t, result, money.Doller(1))
}

func TestReduceMoneyToDifferentCurrency(t *testing.T) {
	bank := money.NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(money.Franc(2), "USD")

	assert.Equal(t, result, money.Doller(1))
}

func TestIdentifyRate(t *testing.T) {
	bank := money.NewBank()

	assert.Equal(t, 1, bank.Rate("USD", "USD"))
}

func TestMixedAddition(t *testing.T) {
	fiveBucks := money.Doller(5)
	tenFrancs := money.Franc(10)
	bank := money.NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(fiveBucks.Plus(tenFrancs), "USD")

	assert.Equal(t, result, money.Doller(10))
}

func TestSumPlusMoney(t *testing.T) {
	fiveBucks := money.Doller(5)
	tenFrancs := money.Franc(10)
	bank := money.NewBank()
	bank.AddRate("CHF", "USD", 2)
	sum := (&money.Sum{fiveBucks, tenFrancs}).Plus(fiveBucks)
	result := bank.Reduce(sum, "USD")

	assert.Equal(t, result, money.Doller(15))
}

func TestSumTimes(t *testing.T) {
	fiveBucks := money.Doller(5)
	tenFrancs := money.Franc(10)
	bank := money.NewBank()
	bank.AddRate("CHF", "USD", 2)
	sum := (&money.Sum{fiveBucks, tenFrancs}).Times(2)
	result := bank.Reduce(sum, "USD")

	assert.Equal(t, result, money.Doller(20))

}

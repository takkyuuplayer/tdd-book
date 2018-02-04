package money_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takkyuuplayer/tdd-book/money"
)

func TestMultipulation(t *testing.T) {
	five := money.NewDoller(5)

	assert.Equal(t, money.NewDoller(10), five.Times(2))
	assert.Equal(t, money.NewDoller(15), five.Times(3))
}

func TestEquality(t *testing.T) {
	assert.Equal(t, money.NewDoller(5).Equals(money.NewDoller(5)), true)
	assert.Equal(t, money.NewDoller(5).Equals(money.NewDoller(6)), false)

	assert.Equal(t, money.NewFranc(5).Equals(money.NewFranc(5)), true)
	assert.Equal(t, money.NewFranc(5).Equals(money.NewFranc(6)), false)

	// assert.Equal(t, money.NewDoller(5).Equals(money.NewFranc(5)), false)
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, money.NewDoller(5).Currency(), "USD")
	assert.Equal(t, money.NewFranc(5).Currency(), "CHF")
}

func TestFrancMultipulation(t *testing.T) {
	five := money.NewFranc(5)

	assert.Equal(t, money.NewFranc(10), five.Times(2))
	assert.Equal(t, money.NewFranc(15), five.Times(3))
}

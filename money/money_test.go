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
	five := money.NewDoller(5)

	assert.Equal(t, five.Equals(money.NewDoller(5)), true)
	assert.Equal(t, five.Equals(money.NewDoller(6)), false)
}

func TestFrancMultipulation(t *testing.T) {
	five := money.NewFranc(5)

	assert.Equal(t, money.NewFranc(10), five.Times(2))
	assert.Equal(t, money.NewFranc(15), five.Times(3))
}

func TestFrancEquality(t *testing.T) {
	five := money.NewFranc(5)

	assert.Equal(t, five.Equals(money.NewFranc(5)), true)
	assert.Equal(t, five.Equals(money.NewFranc(6)), false)
}

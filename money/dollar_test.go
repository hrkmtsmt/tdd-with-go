package money_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"app/money"
)

func TestDollerMultiplication(t *testing.T) {
	five := money.NewDollar(5)

	five.Times(2)
	result := five.Times(3)

	assert.Equal(t, result, money.NewDollar(15))
}

func TestDollarEquality(t *testing.T) {
	assert.True(t, money.NewDollar(5).Equals(money.NewDollar(5)))
	assert.False(t, money.NewDollar(5).Equals(money.NewDollar(6)))
}

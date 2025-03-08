package money_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"app/money"
)

func TestDollerMultiplication(t *testing.T) {
	product := money.NewDollar(5)

	product.Times(2)
	result := product.Times(3)

	assert.Equal(t, result.Amount, 15)
}

func TestDollarEquality(t *testing.T) {
	assert.True(t, money.NewDollar(5).Equals(money.NewDollar(5)))
	assert.False(t, money.NewDollar(5).Equals(money.NewDollar(6)))
}

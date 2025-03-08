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

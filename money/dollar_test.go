package money_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	m "app/money"
)

func TestDollerMultiplication(t *testing.T) {
	five := m.NewDollar(5)

	five.Times(2)

	assert.Equal(t, five.Amount, 10)
}

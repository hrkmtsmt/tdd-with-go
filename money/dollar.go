package money

type Dollar struct {
	Amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{amount}
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.Amount * multiplier)
}

package money

type Money struct {
	amount   int
	Currency string
}

func (m *Money) reduce(b *Bank, to string) *Money {
	rate := b.Rate(m.Currency, to)
	return &Money{
		m.amount / rate,
		to,
	}
}

func (m *Money) Plus(m2 Expression) Expression {
	return &Sum{m, m2}
}

func (m *Money) Times(multiplier int) Expression {
	return &Money{
		m.amount * multiplier,
		m.Currency,
	}
}

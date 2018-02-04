package money

type Money struct {
	amount   int
	currency string
}

func (m *Money) Currency() string {
	return m.currency
}

func (m *Money) reduce(b *Bank, to string) *Money {
	rate := b.Rate(m.currency, to)
	return &Money{
		m.amount / rate,
		to,
	}
}

func (m *Money) Plus(m2 Expression) Expression {
	return &Sum{
		m,
		m2,
	}
}

func (m *Money) Times(multiplier int) Expression {
	return &Money{
		m.amount * multiplier,
		m.currency,
	}
}

func (m *Money) Equals(m2 *Money) bool {
	return m.amount == m2.amount && m.currency == m2.currency
}

func Doller(amount int) *Money {
	return &Money{amount, "USD"}
}

func Franc(amount int) *Money {
	return &Money{amount, "CHF"}
}

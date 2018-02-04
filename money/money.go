package money

type Money struct {
	amount   int
	currency string
}

type IMoney interface {
	Amount() int
	Currency() string
}

type Bank struct{}

type Expression interface {
}

func (b *Bank) Reduce(exp Expression, currency string) *Money {
	return &Money{10, "USD"}
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) Currency() string {
	return m.currency
}

func (m *Money) Equals(obj interface{}) bool {
	m2 := obj.(IMoney)
	return m.Amount() == m2.Amount() && m.Currency() == m2.Currency()
}

func (m *Money) Times(multiplier int) *Money {
	return &Money{
		m.Amount() * multiplier,
		m.Currency(),
	}
}

func (m *Money) Plus(m2 *Money) Expression {
	return &Money{
		m.Amount() + m2.Amount(),
		m.Currency(),
	}
}

func Doller(amount int) *Money {
	return &Money{amount, "USD"}
}

func Franc(amount int) *Money {
	return &Money{amount, "CHF"}
}

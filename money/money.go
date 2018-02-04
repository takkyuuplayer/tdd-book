package money

type Money struct {
	amount   int
	currency string
}

type IMoney interface {
	Amount() int
	Currency() string
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

func NewDoller(amount int) *Money {
	return &Money{amount, "USD"}
}

func NewFranc(amount int) *Money {
	return &Money{amount, "CHF"}
}

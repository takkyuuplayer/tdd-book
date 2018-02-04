package money

type Money struct {
	amount   int
	currency string
}

type IMoney interface {
	Amount() int
	Currency() string
}

type Bank struct {
	rates map[Pair]int
}

type Expression interface {
	reduce(b *Bank, currency string) *Money
	Plus(addend Expression) Expression
}

type Sum struct {
	Augend Expression
	Addend Expression
}

type Pair struct {
	from, to string
}

func (s *Sum) reduce(b *Bank, currency string) *Money {
	amount := s.Augend.reduce(b, currency).amount + s.Addend.reduce(b, currency).amount
	return &Money{
		amount,
		currency,
	}
}

func (s *Sum) Plus(addend Expression) Expression {
	return &Sum{s, addend}
}

func (b *Bank) Reduce(source Expression, currency string) *Money {
	return source.reduce(b, currency)
}
func (b *Bank) AddRate(from, to string, rate int) {
	b.rates[Pair{from, to}] = rate
}

func (b *Bank) Rate(from, to string) int {
	if from == to {
		return 1
	}
	return b.rates[Pair{from, to}]
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) reduce(b *Bank, to string) *Money {
	rate := b.Rate(m.currency, to)
	return &Money{
		m.amount / rate,
		to,
	}
}

func (m *Money) Currency() string {
	return m.currency
}

func (m *Money) Equals(obj interface{}) bool {
	m2 := obj.(IMoney)
	return m.Amount() == m2.Amount() && m.Currency() == m2.Currency()
}

func (m *Money) Times(multiplier int) Expression {
	return &Money{
		m.Amount() * multiplier,
		m.Currency(),
	}
}

func (m *Money) Plus(m2 Expression) Expression {
	return &Sum{
		m,
		m2,
	}
}

func Doller(amount int) *Money {
	return &Money{amount, "USD"}
}

func Franc(amount int) *Money {
	return &Money{amount, "CHF"}
}

func NewBank() *Bank {
	return &Bank{make(map[Pair]int)}
}

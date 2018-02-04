package money

type Pair struct {
	from, to string
}

type Bank struct {
	rates map[Pair]int
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

func NewBank() *Bank {
	return &Bank{make(map[Pair]int)}
}

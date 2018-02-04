package money

type Money struct {
	amount int
}

type IMoney interface {
	Amount() int
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) Equals(obj interface{}) bool {
	m2 := obj.(IMoney)
	return m.Amount() == m2.Amount()
}

type Doller struct {
	Money
}

func (d *Doller) amount() int {
	return d.Money.amount
}

func NewDoller(amount int) *Doller {
	return &Doller{Money{amount: amount}}
}

func (d *Doller) Times(multiplier int) *Doller {
	return NewDoller(d.amount() * multiplier)
}

type Franc struct {
	Money
}

func (f *Franc) amount() int {
	return f.Money.amount
}

func NewFranc(amount int) *Franc {
	return &Franc{Money{amount: amount}}
}

func (d *Franc) Times(multiplier int) *Franc {
	return NewFranc(d.amount() * multiplier)
}

package money

type Doller struct {
	amount int
}

func NewDoller(amount int) *Doller {
	return &Doller{amount: amount}
}

func (d *Doller) Times(multiplier int) *Doller {
	return NewDoller(d.amount * multiplier)
}

func (d *Doller) Equals(obj interface{}) bool {
	doller := obj.(*Doller)
	return d.amount == doller.amount
}

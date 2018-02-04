package money

type Doller struct {
	Amount int
}

func (d *Doller) Times(multiplier int) *Doller {
	return &Doller{d.Amount * multiplier}
}

func (d *Doller) Equals(obj interface{}) bool {
	doller := obj.(*Doller)
	return d.Amount == doller.Amount
}

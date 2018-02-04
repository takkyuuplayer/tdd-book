package money

type Doller struct {
	Amount int
}

func (d *Doller) Times(multiplier int) *Doller {
	return &Doller{d.Amount * multiplier}
}

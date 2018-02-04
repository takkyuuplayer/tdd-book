package money

type Doller struct {
	Amount int
}

func (d *Doller) Times(multiplier int) {
	d.Amount *= multiplier
}

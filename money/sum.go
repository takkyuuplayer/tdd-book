package money

type Sum struct {
	Augend Expression
	Addend Expression
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

func (s *Sum) Times(multiplier int) Expression {
	return &Sum{
		s.Augend.Times(multiplier),
		s.Addend.Times(multiplier),
	}
}

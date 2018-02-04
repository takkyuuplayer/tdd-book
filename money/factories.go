package money

func Doller(amount int) *Money {
	return &Money{amount, "USD"}
}

func Franc(amount int) *Money {
	return &Money{amount, "CHF"}
}

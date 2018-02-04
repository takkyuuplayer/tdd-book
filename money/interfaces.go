package money

type Expression interface {
	reduce(b *Bank, currency string) *Money
	Plus(addend Expression) Expression
	Times(multiplier int) Expression
}

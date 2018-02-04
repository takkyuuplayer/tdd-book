package money

import "testing"

func TestMultipulation(t *testing.T) {
	five := Doller{5}
	product := five.Times(2)

	if product.Amount != 10 {
		t.Errorf(`product.Amount = %#v, want %#v`, product.Amount, 10)
	}

	product = five.Times(3)

	if product.Amount != 15 {
		t.Errorf(`product.Amount = %#v, want %#v`, product.Amount, 15)
	}
}

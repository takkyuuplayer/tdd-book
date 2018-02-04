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

func TestEquality(t *testing.T) {
	five := &Doller{5}
	five2 := &Doller{5}

	if five.Equals(five2) != true {
		t.Errorf(`five.Equals(five2) = %#v, want %#v`, five.Equals(five2), true)
	}

	six := &Doller{6}

	if five.Equals(six) != false {
		t.Errorf(`five.Equals(six) = %#v, want %#v`, five.Equals(six), false)
	}
}

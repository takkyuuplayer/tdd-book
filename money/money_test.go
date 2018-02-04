package money

import "testing"

func TestMultipulation(t *testing.T) {
	five := Doller{5}
	five.Times(2)
	if five.Amount != 10 {
		t.Errorf(`five.Amount = %#v, want %#v`, five.Amount, 10)
	}

}

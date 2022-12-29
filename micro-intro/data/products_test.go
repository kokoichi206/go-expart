package data

import "testing"

func TestChecksValidation(t *testing.T) {

	p := &Product{
		Name:  "Ore",
		Price: 1.00,
		SKU:   "abc-defg-hijkl",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Hello",
		Price: 1,
		SKU:   "a-b-c",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}

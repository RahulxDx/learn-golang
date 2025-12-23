package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
}

func newBill(n string) bill {
	return bill{
		name: n,
		items: map[string]float64{
			"pie":  5.99,
			"cake": 3.99,
		},
	}
}

func (b bill) show() string {
	s := "Bill:\n"
	var t float64

	for k, v := range b.items {
		s += fmt.Sprintf("%s : %.2f\n", k, v)
		t += v
	}

	s += fmt.Sprintf("Total : %.2f", t)
	return s
}

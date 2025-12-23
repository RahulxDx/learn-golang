package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tp    float64
}

func newBill(n string) bill {
	return bill{
		name:  n,
		items: map[string]float64{},
	}
}

func (b *bill) add(i string, p float64) {
	b.items[i] = p
}

func (b *bill) setTip(t float64) {
	b.tp = t
}

func (b bill) show() string {
	s := b.name + "\n"
	var total float64

	for k, v := range b.items {
		s += fmt.Sprintf("%s : %.2f\n", k, v)
		total += v
	}

	s += fmt.Sprintf("tip : %.2f\n", b.tp)
	s += fmt.Sprintf("total : %.2f", total+b.tp)

	return s
}

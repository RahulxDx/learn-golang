package main

import (
	"fmt"
	"math"
)

type calc interface {
	val() float64
}

type sq struct {
	n float64
}

type cir struct {
	r float64
}

func (s sq) val() float64 {
	return s.n * s.n
}

func (c cir) val() float64 {
	return math.Pi * c.r * c.r
}

func show(c calc) {
	fmt.Println("Value:", c.val())
}

func main() {
	a := sq{n: 6}
	b := cir{r: 4}

	show(a)
	show(b)
}

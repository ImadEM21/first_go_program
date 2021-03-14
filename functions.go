package main

import (
	"errors"
)

func sum(a, b int) int {
	return a + b
}

func makeAdder(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func somme(nbs ...int) int {
	var res int
	for _, i := range nbs {
		res = res + i
	}
	return res
}

func divide(a, b int) (int, error) {
	if b != 0 {
		return a / b, nil
	}
	return 0, errors.New("On ne peut pas diviser un nombre par 0")
}

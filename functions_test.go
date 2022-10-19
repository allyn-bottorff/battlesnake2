package main

import "testing"

func TestInPointSlice(t *testing.T) {
	p := Point{X: 0, Y: 1}
	ps := []Point{
		{X: 4, Y: 1},
		{X: 3, Y: 3},
		{X: 2, Y: 5},
		{X: 0, Y: 1},
		{X: 9, Y: 5},
	}

	if !inPointSlice(p, ps) {
		t.Fatal()
	}
}

func TestNotInPointSlice(t *testing.T) {
	p := Point{X: 10, Y: 1}
	ps := []Point{
		{X: 4, Y: 1},
		{X: 3, Y: 3},
		{X: 2, Y: 5},
		{X: 0, Y: 1},
		{X: 9, Y: 5},
	}

	if inPointSlice(p, ps) {
		t.Fatal()
	}
}

func TestEqualPoints(t *testing.T) {
	p1 := Point{X: 3, Y: 4}
	p2 := Point{X: 3, Y: 4}

	var equals bool
	if p1 == p2 {
		equals = true
	} else {
		equals = false
	}

	if !equals {
		t.Fatal()
	}
}

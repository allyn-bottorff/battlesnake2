package main

import "testing"

func TestInCoordSlice(t *testing.T) {
	p := Coord{X: 0, Y: 1}
	ps := []Coord{
		{X: 4, Y: 1},
		{X: 3, Y: 3},
		{X: 2, Y: 5},
		{X: 0, Y: 1},
		{X: 9, Y: 5},
	}

	if !inCoordSlice(p, ps) {
		t.Fatal()
	}
}

func TestNotInCoordSlice(t *testing.T) {
	p := Coord{X: 10, Y: 1}
	ps := []Coord{
		{X: 4, Y: 1},
		{X: 3, Y: 3},
		{X: 2, Y: 5},
		{X: 0, Y: 1},
		{X: 9, Y: 5},
	}

	if inCoordSlice(p, ps) {
		t.Fatal()
	}
}

package main

import (
	"testing"
)

func TestFindDistanceBetweenCoords0(t *testing.T) {
	coord1 := Coord{X: 0, Y: 0}
	coord2 := Coord{X: 0, Y: 0}
	expected := 0
	actual := FindDistanceBetweenCoords(coord1, coord2)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestFindDistanceBetweenCoords1(t *testing.T) {
	coord1 := Coord{X: 0, Y: 0}
	coord2 := Coord{X: 0, Y: 1}
	expected := 1
	actual := FindDistanceBetweenCoords(coord1, coord2)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestFindDistanceBetweenCoords2(t *testing.T) {
	coord1 := Coord{X: 1, Y: 0}
	coord2 := Coord{X: 0, Y: 1}
	expected := 2
	actual := FindDistanceBetweenCoords(coord1, coord2)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestFindDistanceBetweenCoords5(t *testing.T) {
	coord1 := Coord{X: 0, Y: 0}
	coord2 := Coord{X: 0, Y: 5}
	expected := 5
	actual := FindDistanceBetweenCoords(coord1, coord2)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

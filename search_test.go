package main

import (
	"encoding/json"
	"testing"
)

func TestGetNeighbourCoordsNOTWRAPPEDBTMID(t *testing.T) {

	coord := Coord{X: 5, Y: 0}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"a05898b9-87aa-4518-a98a-81b4d0a06f3c","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":365,"board":{"height":11,"width":11,"food":[{"x":9,"y":0},{"x":8,"y":4},{"x":4,"y":7}],"hazards":[],"snakes":[{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1}],"head":{"x":10,"y":3},"length":35,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"15a95e7e-93c1-43fb-8eca-5be495e8c083","name":"ts-snake","health":99,"body":[{"x":2,"y":5},{"x":3,"y":5},{"x":3,"y":4},{"x":3,"y":3},{"x":3,"y":2},{"x":4,"y":2},{"x":5,"y":2},{"x":6,"y":2},{"x":7,"y":2},{"x":7,"y":3},{"x":6,"y":3},{"x":5,"y":3},{"x":4,"y":3},{"x":4,"y":4},{"x":5,"y":4},{"x":5,"y":5},{"x":6,"y":5},{"x":7,"y":5},{"x":7,"y":6},{"x":7,"y":7},{"x":7,"y":8},{"x":7,"y":9},{"x":6,"y":9},{"x":5,"y":9},{"x":4,"y":9},{"x":3,"y":9},{"x":2,"y":9},{"x":1,"y":9},{"x":1,"y":8},{"x":2,"y":8},{"x":3,"y":8},{"x":4,"y":8},{"x":5,"y":8}],"head":{"x":2,"y":5},"length":33,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1}],"head":{"x":10,"y":3},"length":35,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)

	expected := []Coord{
		{X: 6, Y: 0},
		{X: 5, Y: 1},
		{X: 4, Y: 0},
	}

	actual := GetNeighbourCoordsNOTWRAPPED(coord, state)

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[0] != expected[0] && actual[0] != expected[1] && actual[0] != expected[2] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[1] != expected[0] && actual[1] != expected[1] && actual[1] != expected[2] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[2] != expected[0] && actual[2] != expected[1] && actual[2] != expected[2] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetNeighbourCoordsNOTWRAPPEDTPMID(t *testing.T) {

	coord := Coord{X: 5, Y: 10}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"a05898b9-87aa-4518-a98a-81b4d0a06f3c","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":365,"board":{"height":11,"width":11,"food":[{"x":9,"y":0},{"x":8,"y":4},{"x":4,"y":7}],"hazards":[],"snakes":[{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1}],"head":{"x":10,"y":3},"length":35,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"15a95e7e-93c1-43fb-8eca-5be495e8c083","name":"ts-snake","health":99,"body":[{"x":2,"y":5},{"x":3,"y":5},{"x":3,"y":4},{"x":3,"y":3},{"x":3,"y":2},{"x":4,"y":2},{"x":5,"y":2},{"x":6,"y":2},{"x":7,"y":2},{"x":7,"y":3},{"x":6,"y":3},{"x":5,"y":3},{"x":4,"y":3},{"x":4,"y":4},{"x":5,"y":4},{"x":5,"y":5},{"x":6,"y":5},{"x":7,"y":5},{"x":7,"y":6},{"x":7,"y":7},{"x":7,"y":8},{"x":7,"y":9},{"x":6,"y":9},{"x":5,"y":9},{"x":4,"y":9},{"x":3,"y":9},{"x":2,"y":9},{"x":1,"y":9},{"x":1,"y":8},{"x":2,"y":8},{"x":3,"y":8},{"x":4,"y":8},{"x":5,"y":8}],"head":{"x":2,"y":5},"length":33,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1}],"head":{"x":10,"y":3},"length":35,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)

	expected := []Coord{
		{X: 6, Y: 10},
		{X: 5, Y: 9},
		{X: 4, Y: 10},
	}

	actual := GetNeighbourCoordsNOTWRAPPED(coord, state)

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[0] != expected[0] && actual[0] != expected[1] && actual[0] != expected[2] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[1] != expected[0] && actual[1] != expected[1] && actual[1] != expected[2] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[2] != expected[0] && actual[2] != expected[1] && actual[2] != expected[2] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetNeighbourCoordsNOTWRAPPEDLFTMID(t *testing.T) {

	coord := Coord{X: 0, Y: 5}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"a05898b9-87aa-4518-a98a-81b4d0a06f3c","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":365,"board":{"height":11,"width":11,"food":[{"x":9,"y":0},{"x":8,"y":4},{"x":4,"y":7}],"hazards":[],"snakes":[{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1}],"head":{"x":10,"y":3},"length":35,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"15a95e7e-93c1-43fb-8eca-5be495e8c083","name":"ts-snake","health":99,"body":[{"x":2,"y":5},{"x":3,"y":5},{"x":3,"y":4},{"x":3,"y":3},{"x":3,"y":2},{"x":4,"y":2},{"x":5,"y":2},{"x":6,"y":2},{"x":7,"y":2},{"x":7,"y":3},{"x":6,"y":3},{"x":5,"y":3},{"x":4,"y":3},{"x":4,"y":4},{"x":5,"y":4},{"x":5,"y":5},{"x":6,"y":5},{"x":7,"y":5},{"x":7,"y":6},{"x":7,"y":7},{"x":7,"y":8},{"x":7,"y":9},{"x":6,"y":9},{"x":5,"y":9},{"x":4,"y":9},{"x":3,"y":9},{"x":2,"y":9},{"x":1,"y":9},{"x":1,"y":8},{"x":2,"y":8},{"x":3,"y":8},{"x":4,"y":8},{"x":5,"y":8}],"head":{"x":2,"y":5},"length":33,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1}],"head":{"x":10,"y":3},"length":35,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)

	expected := []Coord{
		{X: 0, Y: 6},
		{X: 1, Y: 5},
		{X: 0, Y: 4},
	}

	actual := GetNeighbourCoordsNOTWRAPPED(coord, state)

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[0] != expected[0] && actual[0] != expected[1] && actual[0] != expected[2] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[1] != expected[0] && actual[1] != expected[1] && actual[1] != expected[2] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[2] != expected[0] && actual[2] != expected[1] && actual[2] != expected[2] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetNeighbourCoordsNOTWRAPPEDRTMID(t *testing.T) {

	coord := Coord{X: 10, Y: 5}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"a05898b9-87aa-4518-a98a-81b4d0a06f3c","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":365,"board":{"height":11,"width":11,"food":[{"x":9,"y":0},{"x":8,"y":4},{"x":4,"y":7}],"hazards":[],"snakes":[{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1}],"head":{"x":10,"y":3},"length":35,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"15a95e7e-93c1-43fb-8eca-5be495e8c083","name":"ts-snake","health":99,"body":[{"x":2,"y":5},{"x":3,"y":5},{"x":3,"y":4},{"x":3,"y":3},{"x":3,"y":2},{"x":4,"y":2},{"x":5,"y":2},{"x":6,"y":2},{"x":7,"y":2},{"x":7,"y":3},{"x":6,"y":3},{"x":5,"y":3},{"x":4,"y":3},{"x":4,"y":4},{"x":5,"y":4},{"x":5,"y":5},{"x":6,"y":5},{"x":7,"y":5},{"x":7,"y":6},{"x":7,"y":7},{"x":7,"y":8},{"x":7,"y":9},{"x":6,"y":9},{"x":5,"y":9},{"x":4,"y":9},{"x":3,"y":9},{"x":2,"y":9},{"x":1,"y":9},{"x":1,"y":8},{"x":2,"y":8},{"x":3,"y":8},{"x":4,"y":8},{"x":5,"y":8}],"head":{"x":2,"y":5},"length":33,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1}],"head":{"x":10,"y":3},"length":35,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)

	expected := []Coord{
		{X: 10, Y: 6},
		{X: 9, Y: 5},
		{X: 10, Y: 4},
	}

	actual := GetNeighbourCoordsNOTWRAPPED(coord, state)

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[0] != expected[0] && actual[0] != expected[1] && actual[0] != expected[2] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[1] != expected[0] && actual[1] != expected[1] && actual[1] != expected[2] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[2] != expected[0] && actual[2] != expected[1] && actual[2] != expected[2] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetNeighbourCoordsNOTWRAPPEDBTLFT(t *testing.T) {

	coord := Coord{X: 0, Y: 0}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"a05898b9-87aa-4518-a98a-81b4d0a06f3c","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":365,"board":{"height":11,"width":11,"food":[{"x":9,"y":0},{"x":8,"y":4},{"x":4,"y":7}],"hazards":[],"snakes":[{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1}],"head":{"x":10,"y":3},"length":35,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"15a95e7e-93c1-43fb-8eca-5be495e8c083","name":"ts-snake","health":99,"body":[{"x":2,"y":5},{"x":3,"y":5},{"x":3,"y":4},{"x":3,"y":3},{"x":3,"y":2},{"x":4,"y":2},{"x":5,"y":2},{"x":6,"y":2},{"x":7,"y":2},{"x":7,"y":3},{"x":6,"y":3},{"x":5,"y":3},{"x":4,"y":3},{"x":4,"y":4},{"x":5,"y":4},{"x":5,"y":5},{"x":6,"y":5},{"x":7,"y":5},{"x":7,"y":6},{"x":7,"y":7},{"x":7,"y":8},{"x":7,"y":9},{"x":6,"y":9},{"x":5,"y":9},{"x":4,"y":9},{"x":3,"y":9},{"x":2,"y":9},{"x":1,"y":9},{"x":1,"y":8},{"x":2,"y":8},{"x":3,"y":8},{"x":4,"y":8},{"x":5,"y":8}],"head":{"x":2,"y":5},"length":33,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1}],"head":{"x":10,"y":3},"length":35,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)

	expected := []Coord{
		{X: 1, Y: 0},
		{X: 0, Y: 1},
	}

	actual := GetNeighbourCoordsNOTWRAPPED(coord, state)

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[0] != expected[0] && actual[0] != expected[1] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[1] != expected[0] && actual[1] != expected[1] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetNeighbourCoordsNOTWRAPPEDBTRT(t *testing.T) {

	coord := Coord{X: 10, Y: 0}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"a05898b9-87aa-4518-a98a-81b4d0a06f3c","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":365,"board":{"height":11,"width":11,"food":[{"x":9,"y":0},{"x":8,"y":4},{"x":4,"y":7}],"hazards":[],"snakes":[{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1}],"head":{"x":10,"y":3},"length":35,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"15a95e7e-93c1-43fb-8eca-5be495e8c083","name":"ts-snake","health":99,"body":[{"x":2,"y":5},{"x":3,"y":5},{"x":3,"y":4},{"x":3,"y":3},{"x":3,"y":2},{"x":4,"y":2},{"x":5,"y":2},{"x":6,"y":2},{"x":7,"y":2},{"x":7,"y":3},{"x":6,"y":3},{"x":5,"y":3},{"x":4,"y":3},{"x":4,"y":4},{"x":5,"y":4},{"x":5,"y":5},{"x":6,"y":5},{"x":7,"y":5},{"x":7,"y":6},{"x":7,"y":7},{"x":7,"y":8},{"x":7,"y":9},{"x":6,"y":9},{"x":5,"y":9},{"x":4,"y":9},{"x":3,"y":9},{"x":2,"y":9},{"x":1,"y":9},{"x":1,"y":8},{"x":2,"y":8},{"x":3,"y":8},{"x":4,"y":8},{"x":5,"y":8}],"head":{"x":2,"y":5},"length":33,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1}],"head":{"x":10,"y":3},"length":35,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)

	expected := []Coord{
		{X: 9, Y: 0},
		{X: 10, Y: 1},
	}

	actual := GetNeighbourCoordsNOTWRAPPED(coord, state)

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[0] != expected[0] && actual[0] != expected[1] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[1] != expected[0] && actual[1] != expected[1] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetNeighbourCoordsNOTWRAPPEDTPLFT(t *testing.T) {

	coord := Coord{X: 0, Y: 10}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"a05898b9-87aa-4518-a98a-81b4d0a06f3c","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":365,"board":{"height":11,"width":11,"food":[{"x":9,"y":0},{"x":8,"y":4},{"x":4,"y":7}],"hazards":[],"snakes":[{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1}],"head":{"x":10,"y":3},"length":35,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"15a95e7e-93c1-43fb-8eca-5be495e8c083","name":"ts-snake","health":99,"body":[{"x":2,"y":5},{"x":3,"y":5},{"x":3,"y":4},{"x":3,"y":3},{"x":3,"y":2},{"x":4,"y":2},{"x":5,"y":2},{"x":6,"y":2},{"x":7,"y":2},{"x":7,"y":3},{"x":6,"y":3},{"x":5,"y":3},{"x":4,"y":3},{"x":4,"y":4},{"x":5,"y":4},{"x":5,"y":5},{"x":6,"y":5},{"x":7,"y":5},{"x":7,"y":6},{"x":7,"y":7},{"x":7,"y":8},{"x":7,"y":9},{"x":6,"y":9},{"x":5,"y":9},{"x":4,"y":9},{"x":3,"y":9},{"x":2,"y":9},{"x":1,"y":9},{"x":1,"y":8},{"x":2,"y":8},{"x":3,"y":8},{"x":4,"y":8},{"x":5,"y":8}],"head":{"x":2,"y":5},"length":33,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1}],"head":{"x":10,"y":3},"length":35,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)

	expected := []Coord{
		{X: 1, Y: 10},
		{X: 0, Y: 9},
	}

	actual := GetNeighbourCoordsNOTWRAPPED(coord, state)

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[0] != expected[0] && actual[0] != expected[1] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[1] != expected[0] && actual[1] != expected[1] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetNeighbourCoordsNOTWRAPPEDTPRT(t *testing.T) {

	coord := Coord{X: 10, Y: 10}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"a05898b9-87aa-4518-a98a-81b4d0a06f3c","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":365,"board":{"height":11,"width":11,"food":[{"x":9,"y":0},{"x":8,"y":4},{"x":4,"y":7}],"hazards":[],"snakes":[{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1}],"head":{"x":10,"y":3},"length":35,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"15a95e7e-93c1-43fb-8eca-5be495e8c083","name":"ts-snake","health":99,"body":[{"x":2,"y":5},{"x":3,"y":5},{"x":3,"y":4},{"x":3,"y":3},{"x":3,"y":2},{"x":4,"y":2},{"x":5,"y":2},{"x":6,"y":2},{"x":7,"y":2},{"x":7,"y":3},{"x":6,"y":3},{"x":5,"y":3},{"x":4,"y":3},{"x":4,"y":4},{"x":5,"y":4},{"x":5,"y":5},{"x":6,"y":5},{"x":7,"y":5},{"x":7,"y":6},{"x":7,"y":7},{"x":7,"y":8},{"x":7,"y":9},{"x":6,"y":9},{"x":5,"y":9},{"x":4,"y":9},{"x":3,"y":9},{"x":2,"y":9},{"x":1,"y":9},{"x":1,"y":8},{"x":2,"y":8},{"x":3,"y":8},{"x":4,"y":8},{"x":5,"y":8}],"head":{"x":2,"y":5},"length":33,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1}],"head":{"x":10,"y":3},"length":35,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)

	expected := []Coord{
		{X: 9, Y: 10},
		{X: 10, Y: 9},
	}

	actual := GetNeighbourCoordsNOTWRAPPED(coord, state)

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[0] != expected[0] && actual[0] != expected[1] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if actual[1] != expected[0] && actual[1] != expected[1] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

package main

import (
	"encoding/json"
	"testing"
)

func TestAvoidNeckUP(t *testing.T) {

	scoredMoves := ScoredMoves{}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"a05898b9-87aa-4518-a98a-81b4d0a06f3c","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":365,"board":{"height":11,"width":11,"food":[{"x":9,"y":0},{"x":8,"y":4},{"x":4,"y":7}],"hazards":[],"snakes":[{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1}],"head":{"x":10,"y":3},"length":35,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"15a95e7e-93c1-43fb-8eca-5be495e8c083","name":"ts-snake","health":99,"body":[{"x":2,"y":5},{"x":3,"y":5},{"x":3,"y":4},{"x":3,"y":3},{"x":3,"y":2},{"x":4,"y":2},{"x":5,"y":2},{"x":6,"y":2},{"x":7,"y":2},{"x":7,"y":3},{"x":6,"y":3},{"x":5,"y":3},{"x":4,"y":3},{"x":4,"y":4},{"x":5,"y":4},{"x":5,"y":5},{"x":6,"y":5},{"x":7,"y":5},{"x":7,"y":6},{"x":7,"y":7},{"x":7,"y":8},{"x":7,"y":9},{"x":6,"y":9},{"x":5,"y":9},{"x":4,"y":9},{"x":3,"y":9},{"x":2,"y":9},{"x":1,"y":9},{"x":1,"y":8},{"x":2,"y":8},{"x":3,"y":8},{"x":4,"y":8},{"x":5,"y":8}],"head":{"x":2,"y":5},"length":33,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1}],"head":{"x":10,"y":3},"length":35,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)
	//fmt.Printf("%v", state)
	//fmt.Printf("your head: %v", state.You.Head)
	//fmt.Printf("your neck: %v", state.You.Body[1])
	scoredMoves = AvoidNeck(state, scoredMoves)
	expected := -1000
	actual := scoredMoves.Up.Score
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestAvoidNeckDOWN(t *testing.T) {

	scoredMoves := ScoredMoves{}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"a05898b9-87aa-4518-a98a-81b4d0a06f3c","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":356,"board":{"height":11,"width":11,"food":[{"x":9,"y":0},{"x":3,"y":4},{"x":8,"y":4},{"x":4,"y":7},{"x":3,"y":5},{"x":5,"y":2}],"hazards":[],"snakes":[{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1},{"x":0,"y":0},{"x":1,"y":0},{"x":2,"y":0},{"x":3,"y":0},{"x":4,"y":0},{"x":5,"y":0},{"x":6,"y":0},{"x":6,"y":1}],"head":{"x":9,"y":7},"length":34,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"15a95e7e-93c1-43fb-8eca-5be495e8c083","name":"ts-snake","health":99,"body":[{"x":7,"y":3},{"x":6,"y":3},{"x":5,"y":3},{"x":4,"y":3},{"x":4,"y":4},{"x":5,"y":4},{"x":5,"y":5},{"x":6,"y":5},{"x":7,"y":5},{"x":7,"y":6},{"x":7,"y":7},{"x":7,"y":8},{"x":7,"y":9},{"x":6,"y":9},{"x":5,"y":9},{"x":4,"y":9},{"x":3,"y":9},{"x":2,"y":9},{"x":1,"y":9},{"x":1,"y":8},{"x":2,"y":8},{"x":3,"y":8},{"x":4,"y":8},{"x":5,"y":8},{"x":6,"y":8},{"x":6,"y":7},{"x":6,"y":6},{"x":5,"y":6},{"x":4,"y":6},{"x":3,"y":6}],"head":{"x":7,"y":3},"length":30,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":97,"body":[{"x":9,"y":7},{"x":9,"y":6},{"x":9,"y":5},{"x":8,"y":5},{"x":8,"y":6},{"x":8,"y":7},{"x":8,"y":8},{"x":8,"y":9},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1},{"x":0,"y":0},{"x":1,"y":0},{"x":2,"y":0},{"x":3,"y":0},{"x":4,"y":0},{"x":5,"y":0},{"x":6,"y":0},{"x":6,"y":1}],"head":{"x":9,"y":7},"length":34,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)
	//fmt.Printf("%v", state)
	//fmt.Printf("your head: %v", state.You.Head)
	//fmt.Printf("your neck: %v", state.You.Body[1])
	scoredMoves = AvoidNeck(state, scoredMoves)
	expected := -1000
	actual := scoredMoves.Down.Score
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestAvoidNeckLEFT(t *testing.T) {

	scoredMoves := ScoredMoves{}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"a05898b9-87aa-4518-a98a-81b4d0a06f3c","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":347,"board":{"height":11,"width":11,"food":[{"x":9,"y":0},{"x":8,"y":9},{"x":3,"y":4},{"x":4,"y":4},{"x":6,"y":3},{"x":8,"y":4},{"x":8,"y":5},{"x":5,"y":3}],"hazards":[],"snakes":[{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":94,"body":[{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1},{"x":0,"y":0},{"x":1,"y":0},{"x":2,"y":0},{"x":3,"y":0},{"x":4,"y":0},{"x":5,"y":0},{"x":6,"y":0},{"x":6,"y":1},{"x":7,"y":1},{"x":8,"y":1},{"x":9,"y":1},{"x":10,"y":1},{"x":10,"y":2},{"x":10,"y":3},{"x":10,"y":4}],"head":{"x":7,"y":10},"length":32,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"15a95e7e-93c1-43fb-8eca-5be495e8c083","name":"ts-snake","health":97,"body":[{"x":7,"y":6},{"x":7,"y":7},{"x":7,"y":8},{"x":7,"y":9},{"x":6,"y":9},{"x":5,"y":9},{"x":4,"y":9},{"x":3,"y":9},{"x":2,"y":9},{"x":1,"y":9},{"x":1,"y":8},{"x":2,"y":8},{"x":3,"y":8},{"x":4,"y":8},{"x":5,"y":8},{"x":6,"y":8},{"x":6,"y":7},{"x":6,"y":6},{"x":5,"y":6},{"x":4,"y":6},{"x":3,"y":6},{"x":3,"y":7},{"x":2,"y":7},{"x":2,"y":6},{"x":1,"y":6},{"x":1,"y":5},{"x":1,"y":4}],"head":{"x":7,"y":6},"length":27,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":94,"body":[{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":10},{"x":0,"y":9},{"x":0,"y":8},{"x":0,"y":7},{"x":0,"y":6},{"x":0,"y":5},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1},{"x":0,"y":0},{"x":1,"y":0},{"x":2,"y":0},{"x":3,"y":0},{"x":4,"y":0},{"x":5,"y":0},{"x":6,"y":0},{"x":6,"y":1},{"x":7,"y":1},{"x":8,"y":1},{"x":9,"y":1},{"x":10,"y":1},{"x":10,"y":2},{"x":10,"y":3},{"x":10,"y":4}],"head":{"x":7,"y":10},"length":32,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)
	//fmt.Printf("%v", state)
	//fmt.Printf("your head: %v", state.You.Head)
	//fmt.Printf("your neck: %v", state.You.Body[1])
	scoredMoves = AvoidNeck(state, scoredMoves)
	expected := -1000
	actual := scoredMoves.Left.Score
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestAvoidNeckRIGHT(t *testing.T) {

	scoredMoves := ScoredMoves{}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"a05898b9-87aa-4518-a98a-81b4d0a06f3c","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":326,"board":{"height":11,"width":11,"food":[{"x":0,"y":2},{"x":0,"y":10},{"x":7,"y":9},{"x":9,"y":0},{"x":0,"y":3},{"x":0,"y":1},{"x":8,"y":9},{"x":0,"y":9},{"x":3,"y":4},{"x":0,"y":7},{"x":4,"y":4},{"x":6,"y":3},{"x":8,"y":4},{"x":3,"y":0},{"x":2,"y":9}],"hazards":[],"snakes":[{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":100,"body":[{"x":4,"y":0},{"x":5,"y":0},{"x":6,"y":0},{"x":6,"y":1},{"x":7,"y":1},{"x":8,"y":1},{"x":9,"y":1},{"x":10,"y":1},{"x":10,"y":2},{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":10},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":4,"y":10}],"head":{"x":4,"y":0},"length":24,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"15a95e7e-93c1-43fb-8eca-5be495e8c083","name":"ts-snake","health":89,"body":[{"x":3,"y":7},{"x":2,"y":7},{"x":2,"y":6},{"x":1,"y":6},{"x":1,"y":5},{"x":1,"y":4},{"x":1,"y":3},{"x":1,"y":2},{"x":1,"y":1},{"x":2,"y":1},{"x":3,"y":1},{"x":4,"y":1},{"x":4,"y":2},{"x":5,"y":2},{"x":6,"y":2},{"x":7,"y":2},{"x":8,"y":2},{"x":8,"y":3},{"x":9,"y":3},{"x":9,"y":4},{"x":9,"y":5},{"x":9,"y":6},{"x":8,"y":6},{"x":8,"y":7},{"x":7,"y":7}],"head":{"x":3,"y":7},"length":25,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":100,"body":[{"x":4,"y":0},{"x":5,"y":0},{"x":6,"y":0},{"x":6,"y":1},{"x":7,"y":1},{"x":8,"y":1},{"x":9,"y":1},{"x":10,"y":1},{"x":10,"y":2},{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":10},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":4,"y":10}],"head":{"x":4,"y":0},"length":24,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)
	//fmt.Printf("%v", state)
	//fmt.Printf("your head: %v", state.You.Head)
	//fmt.Printf("your neck: %v", state.You.Body[1])
	scoredMoves = AvoidNeck(state, scoredMoves)
	expected := -1000
	actual := scoredMoves.Right.Score
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestAvoidWallsDOWN(t *testing.T) {
	scoredMoves := ScoredMoves{}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"a05898b9-87aa-4518-a98a-81b4d0a06f3c","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":326,"board":{"height":11,"width":11,"food":[{"x":0,"y":2},{"x":0,"y":10},{"x":7,"y":9},{"x":9,"y":0},{"x":0,"y":3},{"x":0,"y":1},{"x":8,"y":9},{"x":0,"y":9},{"x":3,"y":4},{"x":0,"y":7},{"x":4,"y":4},{"x":6,"y":3},{"x":8,"y":4},{"x":3,"y":0},{"x":2,"y":9}],"hazards":[],"snakes":[{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":100,"body":[{"x":4,"y":0},{"x":5,"y":0},{"x":6,"y":0},{"x":6,"y":1},{"x":7,"y":1},{"x":8,"y":1},{"x":9,"y":1},{"x":10,"y":1},{"x":10,"y":2},{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":10},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":4,"y":10}],"head":{"x":4,"y":0},"length":24,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"15a95e7e-93c1-43fb-8eca-5be495e8c083","name":"ts-snake","health":89,"body":[{"x":3,"y":7},{"x":2,"y":7},{"x":2,"y":6},{"x":1,"y":6},{"x":1,"y":5},{"x":1,"y":4},{"x":1,"y":3},{"x":1,"y":2},{"x":1,"y":1},{"x":2,"y":1},{"x":3,"y":1},{"x":4,"y":1},{"x":4,"y":2},{"x":5,"y":2},{"x":6,"y":2},{"x":7,"y":2},{"x":8,"y":2},{"x":8,"y":3},{"x":9,"y":3},{"x":9,"y":4},{"x":9,"y":5},{"x":9,"y":6},{"x":8,"y":6},{"x":8,"y":7},{"x":7,"y":7}],"head":{"x":3,"y":7},"length":25,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":100,"body":[{"x":4,"y":0},{"x":5,"y":0},{"x":6,"y":0},{"x":6,"y":1},{"x":7,"y":1},{"x":8,"y":1},{"x":9,"y":1},{"x":10,"y":1},{"x":10,"y":2},{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":10},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":4,"y":10}],"head":{"x":4,"y":0},"length":24,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)
	//fmt.Printf("%v", state)
	//fmt.Printf("your head: %v", state.You.Head)
	//fmt.Printf("your neck: %v", state.You.Body[1])
	scoredMoves = AvoidWalls(state, scoredMoves)
	expected := -1000
	actual := scoredMoves.Down.Score
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestAvoidWallsRIGHT(t *testing.T) {
	scoredMoves := ScoredMoves{}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"a05898b9-87aa-4518-a98a-81b4d0a06f3c","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":317,"board":{"height":11,"width":11,"food":[{"x":4,"y":0},{"x":6,"y":0},{"x":7,"y":1},{"x":0,"y":2},{"x":0,"y":10},{"x":5,"y":0},{"x":7,"y":9},{"x":9,"y":0},{"x":0,"y":3},{"x":0,"y":1},{"x":8,"y":9},{"x":0,"y":9},{"x":3,"y":4},{"x":0,"y":7},{"x":10,"y":1},{"x":4,"y":4},{"x":8,"y":1},{"x":6,"y":3},{"x":8,"y":4},{"x":3,"y":0}],"hazards":[],"snakes":[{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":87,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":10},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":2,"y":9},{"x":3,"y":9}],"head":{"x":10,"y":3},"length":18,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"15a95e7e-93c1-43fb-8eca-5be495e8c083","name":"ts-snake","health":98,"body":[{"x":2,"y":1},{"x":3,"y":1},{"x":4,"y":1},{"x":4,"y":2},{"x":5,"y":2},{"x":6,"y":2},{"x":7,"y":2},{"x":8,"y":2},{"x":8,"y":3},{"x":9,"y":3},{"x":9,"y":4},{"x":9,"y":5},{"x":9,"y":6},{"x":8,"y":6},{"x":8,"y":7},{"x":7,"y":7},{"x":7,"y":6},{"x":7,"y":5},{"x":6,"y":5},{"x":5,"y":5},{"x":4,"y":5},{"x":3,"y":5},{"x":2,"y":5},{"x":1,"y":5},{"x":1,"y":4}],"head":{"x":2,"y":1},"length":25,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":87,"body":[{"x":10,"y":3},{"x":10,"y":4},{"x":10,"y":5},{"x":10,"y":6},{"x":10,"y":7},{"x":10,"y":8},{"x":10,"y":9},{"x":9,"y":9},{"x":9,"y":10},{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":2,"y":9},{"x":3,"y":9}],"head":{"x":10,"y":3},"length":18,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)
	//fmt.Printf("%v", state)
	//fmt.Printf("your head: %v", state.You.Head)
	//fmt.Printf("your neck: %v", state.You.Body[1])
	scoredMoves = AvoidWalls(state, scoredMoves)
	expected := -1000
	actual := scoredMoves.Right.Score
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestAvoidWallsUP(t *testing.T) {
	scoredMoves := ScoredMoves{}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"a05898b9-87aa-4518-a98a-81b4d0a06f3c","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":308,"board":{"height":11,"width":11,"food":[{"x":4,"y":0},{"x":6,"y":0},{"x":7,"y":1},{"x":0,"y":2},{"x":0,"y":10},{"x":5,"y":0},{"x":7,"y":9},{"x":9,"y":0},{"x":0,"y":3},{"x":0,"y":1},{"x":8,"y":9},{"x":0,"y":9},{"x":3,"y":4},{"x":0,"y":7},{"x":4,"y":1},{"x":10,"y":1},{"x":4,"y":4},{"x":8,"y":1},{"x":6,"y":3},{"x":8,"y":4}],"hazards":[],"snakes":[{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":96,"body":[{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":2,"y":9},{"x":3,"y":9},{"x":4,"y":9},{"x":5,"y":9},{"x":6,"y":9},{"x":6,"y":8},{"x":5,"y":8},{"x":4,"y":8},{"x":3,"y":8},{"x":2,"y":8},{"x":2,"y":7}],"head":{"x":8,"y":10},"length":18,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"15a95e7e-93c1-43fb-8eca-5be495e8c083","name":"ts-snake","health":92,"body":[{"x":9,"y":3},{"x":9,"y":4},{"x":9,"y":5},{"x":9,"y":6},{"x":8,"y":6},{"x":8,"y":7},{"x":7,"y":7},{"x":7,"y":6},{"x":7,"y":5},{"x":6,"y":5},{"x":5,"y":5},{"x":4,"y":5},{"x":3,"y":5},{"x":2,"y":5},{"x":1,"y":5},{"x":1,"y":4},{"x":1,"y":3},{"x":1,"y":2},{"x":2,"y":2},{"x":3,"y":2},{"x":4,"y":2},{"x":5,"y":2},{"x":6,"y":2},{"x":7,"y":2}],"head":{"x":9,"y":3},"length":24,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"893799d4-c6df-4bec-8188-d968a1d9a9a3","name":"Go Starter Project","health":96,"body":[{"x":8,"y":10},{"x":7,"y":10},{"x":6,"y":10},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":2,"y":9},{"x":3,"y":9},{"x":4,"y":9},{"x":5,"y":9},{"x":6,"y":9},{"x":6,"y":8},{"x":5,"y":8},{"x":4,"y":8},{"x":3,"y":8},{"x":2,"y":8},{"x":2,"y":7}],"head":{"x":8,"y":10},"length":18,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)
	//fmt.Printf("%v", state)
	//fmt.Printf("your head: %v", state.You.Head)
	//fmt.Printf("your neck: %v", state.You.Body[1])
	scoredMoves = AvoidWalls(state, scoredMoves)
	expected := -1000
	actual := scoredMoves.Up.Score
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestAvoidWallsLEFT(t *testing.T) {
	scoredMoves := ScoredMoves{}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"3800233a-2f09-40a3-8fba-e682ec327cdb","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":361,"board":{"height":11,"width":11,"food":[{"x":10,"y":1},{"x":10,"y":0},{"x":5,"y":0},{"x":7,"y":0}],"hazards":[],"snakes":[{"id":"cb33fb86-8865-4230-8f7c-42f6d1782fec","name":"Go Starter Project","health":99,"body":[{"x":0,"y":3},{"x":0,"y":4},{"x":0,"y":5},{"x":0,"y":6},{"x":0,"y":7},{"x":0,"y":8},{"x":0,"y":9},{"x":0,"y":10},{"x":1,"y":10},{"x":1,"y":9},{"x":1,"y":8},{"x":1,"y":7},{"x":1,"y":6},{"x":1,"y":5},{"x":1,"y":4},{"x":1,"y":3},{"x":2,"y":3},{"x":2,"y":4},{"x":2,"y":5},{"x":2,"y":6},{"x":2,"y":7},{"x":2,"y":8},{"x":2,"y":9},{"x":2,"y":10}],"head":{"x":0,"y":3},"length":24,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"5c53e05d-f748-41fe-ab9b-d30e4b7de7ec","name":"ts-snake","health":88,"body":[{"x":9,"y":6},{"x":8,"y":6},{"x":8,"y":7},{"x":7,"y":7},{"x":7,"y":8},{"x":6,"y":8},{"x":5,"y":8},{"x":5,"y":7},{"x":6,"y":7},{"x":6,"y":6},{"x":5,"y":6},{"x":5,"y":5},{"x":5,"y":4},{"x":5,"y":3},{"x":5,"y":2},{"x":4,"y":2},{"x":3,"y":2},{"x":3,"y":1},{"x":4,"y":1},{"x":5,"y":1},{"x":6,"y":1},{"x":7,"y":1},{"x":8,"y":1},{"x":9,"y":1}],"head":{"x":9,"y":6},"length":24,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"cb33fb86-8865-4230-8f7c-42f6d1782fec","name":"Go Starter Project","health":99,"body":[{"x":0,"y":3},{"x":0,"y":4},{"x":0,"y":5},{"x":0,"y":6},{"x":0,"y":7},{"x":0,"y":8},{"x":0,"y":9},{"x":0,"y":10},{"x":1,"y":10},{"x":1,"y":9},{"x":1,"y":8},{"x":1,"y":7},{"x":1,"y":6},{"x":1,"y":5},{"x":1,"y":4},{"x":1,"y":3},{"x":2,"y":3},{"x":2,"y":4},{"x":2,"y":5},{"x":2,"y":6},{"x":2,"y":7},{"x":2,"y":8},{"x":2,"y":9},{"x":2,"y":10}],"head":{"x":0,"y":3},"length":24,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)
	//fmt.Printf("%v", state)
	//fmt.Printf("your head: %v", state.You.Head)
	//fmt.Printf("your neck: %v", state.You.Body[1])
	scoredMoves = AvoidWalls(state, scoredMoves)
	expected := -1000
	actual := scoredMoves.Left.Score
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestAvoidSelfDOWNRIGHT(t *testing.T) {
	scoredMoves := ScoredMoves{}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"3800233a-2f09-40a3-8fba-e682ec327cdb","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":352,"board":{"height":11,"width":11,"food":[{"x":0,"y":4},{"x":0,"y":10},{"x":0,"y":5},{"x":10,"y":1},{"x":10,"y":0},{"x":5,"y":0}],"hazards":[],"snakes":[{"id":"cb33fb86-8865-4230-8f7c-42f6d1782fec","name":"Go Starter Project","health":81,"body":[{"x":1,"y":9},{"x":1,"y":8},{"x":1,"y":7},{"x":1,"y":6},{"x":1,"y":5},{"x":1,"y":4},{"x":1,"y":3},{"x":2,"y":3},{"x":2,"y":4},{"x":2,"y":5},{"x":2,"y":6},{"x":2,"y":7},{"x":2,"y":8},{"x":2,"y":9},{"x":2,"y":10},{"x":3,"y":10},{"x":4,"y":10},{"x":5,"y":10},{"x":6,"y":10},{"x":7,"y":10},{"x":8,"y":10}],"head":{"x":1,"y":9},"length":21,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"5c53e05d-f748-41fe-ab9b-d30e4b7de7ec","name":"ts-snake","health":97,"body":[{"x":6,"y":6},{"x":5,"y":6},{"x":5,"y":5},{"x":5,"y":4},{"x":5,"y":3},{"x":5,"y":2},{"x":4,"y":2},{"x":3,"y":2},{"x":3,"y":1},{"x":4,"y":1},{"x":5,"y":1},{"x":6,"y":1},{"x":7,"y":1},{"x":8,"y":1},{"x":9,"y":1},{"x":9,"y":2},{"x":9,"y":3},{"x":9,"y":4},{"x":9,"y":5},{"x":9,"y":6},{"x":9,"y":7},{"x":9,"y":8},{"x":8,"y":8},{"x":7,"y":8}],"head":{"x":6,"y":6},"length":24,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"cb33fb86-8865-4230-8f7c-42f6d1782fec","name":"Go Starter Project","health":81,"body":[{"x":1,"y":9},{"x":1,"y":8},{"x":1,"y":7},{"x":1,"y":6},{"x":1,"y":5},{"x":1,"y":4},{"x":1,"y":3},{"x":2,"y":3},{"x":2,"y":4},{"x":2,"y":5},{"x":2,"y":6},{"x":2,"y":7},{"x":2,"y":8},{"x":2,"y":9},{"x":2,"y":10},{"x":3,"y":10},{"x":4,"y":10},{"x":5,"y":10},{"x":6,"y":10},{"x":7,"y":10},{"x":8,"y":10}],"head":{"x":1,"y":9},"length":21,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)
	//fmt.Printf("%v", state)
	//fmt.Printf("your head: %v", state.You.Head)
	//fmt.Printf("your neck: %v", state.You.Body[1])
	scoredMoves = AvoidSelf(state, scoredMoves)
	expected := -1000
	actualDown := scoredMoves.Down.Score
	actualRight := scoredMoves.Right.Score
	if actualDown != expected || actualRight != expected {
		t.Errorf("Expected Down: %v, got %v", expected, actualDown)
		t.Errorf("Expected Right: %v, got %v", expected, actualRight)
	}
}

func TestAvoidSelfLeft(t *testing.T) {
	scoredMoves := ScoredMoves{}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"3800233a-2f09-40a3-8fba-e682ec327cdb","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":329,"board":{"height":11,"width":11,"food":[{"x":0,"y":4},{"x":0,"y":10},{"x":0,"y":5},{"x":10,"y":1},{"x":10,"y":0},{"x":7,"y":10}],"hazards":[],"snakes":[{"id":"cb33fb86-8865-4230-8f7c-42f6d1782fec","name":"Go Starter Project","health":99,"body":[{"x":8,"y":9},{"x":7,"y":9},{"x":6,"y":9},{"x":5,"y":9},{"x":4,"y":9},{"x":3,"y":9},{"x":3,"y":8},{"x":3,"y":7},{"x":3,"y":6},{"x":3,"y":5},{"x":3,"y":4},{"x":2,"y":4},{"x":2,"y":5},{"x":2,"y":6},{"x":2,"y":7},{"x":2,"y":8},{"x":2,"y":9},{"x":1,"y":9},{"x":1,"y":8},{"x":1,"y":7}],"head":{"x":8,"y":9},"length":20,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"5c53e05d-f748-41fe-ab9b-d30e4b7de7ec","name":"ts-snake","health":96,"body":[{"x":7,"y":8},{"x":6,"y":8},{"x":5,"y":8},{"x":4,"y":8},{"x":4,"y":7},{"x":4,"y":6},{"x":4,"y":5},{"x":4,"y":4},{"x":4,"y":3},{"x":4,"y":2},{"x":5,"y":2},{"x":6,"y":2},{"x":7,"y":2},{"x":8,"y":2},{"x":8,"y":3},{"x":8,"y":4},{"x":8,"y":5},{"x":7,"y":5},{"x":7,"y":4},{"x":7,"y":3},{"x":6,"y":3},{"x":6,"y":4},{"x":6,"y":5}],"head":{"x":7,"y":8},"length":23,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"cb33fb86-8865-4230-8f7c-42f6d1782fec","name":"Go Starter Project","health":99,"body":[{"x":8,"y":9},{"x":7,"y":9},{"x":6,"y":9},{"x":5,"y":9},{"x":4,"y":9},{"x":3,"y":9},{"x":3,"y":8},{"x":3,"y":7},{"x":3,"y":6},{"x":3,"y":5},{"x":3,"y":4},{"x":2,"y":4},{"x":2,"y":5},{"x":2,"y":6},{"x":2,"y":7},{"x":2,"y":8},{"x":2,"y":9},{"x":1,"y":9},{"x":1,"y":8},{"x":1,"y":7}],"head":{"x":8,"y":9},"length":20,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)
	//fmt.Printf("%v", state)
	//fmt.Printf("your head: %v", state.You.Head)
	//fmt.Printf("your neck: %v", state.You.Body[1])
	scoredMoves = AvoidSelf(state, scoredMoves)
	expected := -1000
	actual := scoredMoves.Left.Score

	if actual != expected {
		t.Errorf("Expected: %v, got %v", expected, actual)

	}
}

func TestAvoidSelfUP(t *testing.T) {
	scoredMoves := ScoredMoves{}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"3800233a-2f09-40a3-8fba-e682ec327cdb","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":317,"board":{"height":11,"width":11,"food":[{"x":0,"y":4},{"x":0,"y":10},{"x":0,"y":5},{"x":10,"y":1},{"x":10,"y":0},{"x":4,"y":7},{"x":7,"y":10},{"x":5,"y":2}],"hazards":[],"snakes":[{"id":"cb33fb86-8865-4230-8f7c-42f6d1782fec","name":"Go Starter Project","health":57,"body":[{"x":2,"y":5},{"x":2,"y":6},{"x":2,"y":7},{"x":2,"y":8},{"x":2,"y":9},{"x":1,"y":9},{"x":1,"y":8},{"x":1,"y":7},{"x":1,"y":6},{"x":1,"y":5},{"x":1,"y":4},{"x":1,"y":3},{"x":1,"y":2},{"x":1,"y":1},{"x":1,"y":0},{"x":2,"y":0},{"x":3,"y":0},{"x":4,"y":0},{"x":4,"y":1}],"head":{"x":2,"y":5},"length":19,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"5c53e05d-f748-41fe-ab9b-d30e4b7de7ec","name":"ts-snake","health":95,"body":[{"x":7,"y":2},{"x":8,"y":2},{"x":8,"y":3},{"x":8,"y":4},{"x":8,"y":5},{"x":7,"y":5},{"x":7,"y":4},{"x":7,"y":3},{"x":6,"y":3},{"x":6,"y":4},{"x":6,"y":5},{"x":5,"y":5},{"x":5,"y":6},{"x":5,"y":7},{"x":5,"y":8},{"x":6,"y":8},{"x":6,"y":9},{"x":7,"y":9},{"x":8,"y":9},{"x":9,"y":9},{"x":9,"y":8}],"head":{"x":7,"y":2},"length":21,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"cb33fb86-8865-4230-8f7c-42f6d1782fec","name":"Go Starter Project","health":57,"body":[{"x":2,"y":5},{"x":2,"y":6},{"x":2,"y":7},{"x":2,"y":8},{"x":2,"y":9},{"x":1,"y":9},{"x":1,"y":8},{"x":1,"y":7},{"x":1,"y":6},{"x":1,"y":5},{"x":1,"y":4},{"x":1,"y":3},{"x":1,"y":2},{"x":1,"y":1},{"x":1,"y":0},{"x":2,"y":0},{"x":3,"y":0},{"x":4,"y":0},{"x":4,"y":1}],"head":{"x":2,"y":5},"length":19,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)
	//fmt.Printf("%v", state)
	//fmt.Printf("your head: %v", state.You.Head)
	//fmt.Printf("your neck: %v", state.You.Body[1])
	scoredMoves = AvoidSelf(state, scoredMoves)
	expected := -1000
	actual := scoredMoves.Up.Score

	if actual != expected {
		t.Errorf("Expected: %v, got %v", expected, actual)

	}
}

func TestAvoidSnakeRight(t *testing.T) {
	scoredMoves := ScoredMoves{}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"3800233a-2f09-40a3-8fba-e682ec327cdb","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":280,"board":{"height":11,"width":11,"food":[{"x":0,"y":4},{"x":9,"y":0},{"x":7,"y":0},{"x":0,"y":10},{"x":0,"y":5},{"x":10,"y":1},{"x":10,"y":0},{"x":4,"y":7},{"x":7,"y":5}],"hazards":[],"snakes":[{"id":"cb33fb86-8865-4230-8f7c-42f6d1782fec","name":"Go Starter Project","health":94,"body":[{"x":5,"y":7},{"x":5,"y":8},{"x":5,"y":9},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":2,"y":9},{"x":2,"y":8},{"x":2,"y":7},{"x":2,"y":6},{"x":2,"y":5},{"x":2,"y":4},{"x":2,"y":3},{"x":2,"y":2},{"x":3,"y":2},{"x":3,"y":3},{"x":3,"y":4},{"x":3,"y":5}],"head":{"x":5,"y":7},"length":19,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"5c53e05d-f748-41fe-ab9b-d30e4b7de7ec","name":"ts-snake","health":99,"body":[{"x":6,"y":6},{"x":6,"y":7},{"x":6,"y":8},{"x":7,"y":8},{"x":7,"y":7},{"x":7,"y":6},{"x":8,"y":6},{"x":8,"y":5},{"x":8,"y":4},{"x":8,"y":3},{"x":8,"y":2},{"x":8,"y":1},{"x":7,"y":1},{"x":6,"y":1},{"x":5,"y":1},{"x":5,"y":2},{"x":6,"y":2}],"head":{"x":6,"y":6},"length":17,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"cb33fb86-8865-4230-8f7c-42f6d1782fec","name":"Go Starter Project","health":94,"body":[{"x":5,"y":7},{"x":5,"y":8},{"x":5,"y":9},{"x":5,"y":10},{"x":4,"y":10},{"x":3,"y":10},{"x":2,"y":10},{"x":2,"y":9},{"x":2,"y":8},{"x":2,"y":7},{"x":2,"y":6},{"x":2,"y":5},{"x":2,"y":4},{"x":2,"y":3},{"x":2,"y":2},{"x":3,"y":2},{"x":3,"y":3},{"x":3,"y":4},{"x":3,"y":5}],"head":{"x":5,"y":7},"length":19,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)
	//fmt.Printf("%v", state)
	//fmt.Printf("your head: %v", state.You.Head)
	//fmt.Printf("your neck: %v", state.You.Body[1])
	scoredMoves = AvoidSnakes(state, scoredMoves)
	expected := -1000
	actual := scoredMoves.Right.Score

	if actual != expected {
		t.Errorf("Expected: %v, got %v", expected, actual)

	}
}

func TestAvoidSnakeUp(t *testing.T) {
	scoredMoves := ScoredMoves{}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"b179006e-e165-43a5-ac39-614dd4ec9a7e","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":381,"board":{"height":11,"width":11,"food":[{"x":10,"y":10},{"x":0,"y":10},{"x":3,"y":3},{"x":7,"y":8},{"x":5,"y":2}],"hazards":[],"snakes":[{"id":"fcb38f74-3921-4982-a760-c948d2185a27","name":"Go Starter Project","health":91,"body":[{"x":6,"y":5},{"x":6,"y":4},{"x":6,"y":3},{"x":5,"y":3},{"x":4,"y":3},{"x":4,"y":4},{"x":3,"y":4},{"x":2,"y":4},{"x":1,"y":4},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1},{"x":0,"y":0},{"x":1,"y":0},{"x":2,"y":0},{"x":3,"y":0},{"x":4,"y":0},{"x":5,"y":0},{"x":6,"y":0},{"x":7,"y":0},{"x":8,"y":0},{"x":9,"y":0},{"x":10,"y":0}],"head":{"x":6,"y":5},"length":24,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"8346a240-506b-4117-a3e9-1e114bc287c5","name":"ts-snake","health":98,"body":[{"x":2,"y":5},{"x":1,"y":5},{"x":0,"y":5},{"x":0,"y":6},{"x":0,"y":7},{"x":0,"y":8},{"x":0,"y":9},{"x":1,"y":9},{"x":1,"y":10},{"x":2,"y":10},{"x":3,"y":10},{"x":4,"y":10},{"x":5,"y":10},{"x":6,"y":10},{"x":7,"y":10},{"x":8,"y":10},{"x":9,"y":10},{"x":9,"y":9},{"x":9,"y":8},{"x":9,"y":7},{"x":10,"y":7},{"x":10,"y":6},{"x":10,"y":5},{"x":10,"y":4},{"x":10,"y":3},{"x":9,"y":3},{"x":9,"y":4},{"x":8,"y":4},{"x":8,"y":5},{"x":8,"y":6},{"x":7,"y":6},{"x":6,"y":6},{"x":5,"y":6},{"x":4,"y":6},{"x":3,"y":6},{"x":2,"y":6},{"x":1,"y":6},{"x":1,"y":7},{"x":2,"y":7},{"x":2,"y":8},{"x":2,"y":9},{"x":3,"y":9},{"x":4,"y":9}],"head":{"x":2,"y":5},"length":43,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"fcb38f74-3921-4982-a760-c948d2185a27","name":"Go Starter Project","health":91,"body":[{"x":6,"y":5},{"x":6,"y":4},{"x":6,"y":3},{"x":5,"y":3},{"x":4,"y":3},{"x":4,"y":4},{"x":3,"y":4},{"x":2,"y":4},{"x":1,"y":4},{"x":0,"y":4},{"x":0,"y":3},{"x":0,"y":2},{"x":0,"y":1},{"x":0,"y":0},{"x":1,"y":0},{"x":2,"y":0},{"x":3,"y":0},{"x":4,"y":0},{"x":5,"y":0},{"x":6,"y":0},{"x":7,"y":0},{"x":8,"y":0},{"x":9,"y":0},{"x":10,"y":0}],"head":{"x":6,"y":5},"length":24,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)
	//fmt.Printf("%v", state)
	//fmt.Printf("your head: %v", state.You.Head)
	//fmt.Printf("your neck: %v", state.You.Body[1])
	scoredMoves = AvoidSnakes(state, scoredMoves)
	expected := -1000
	actual := scoredMoves.Up.Score

	if actual != expected {
		t.Errorf("Expected: %v, got %v", expected, actual)

	}
}

func TestAvoidSnakeDown(t *testing.T) {
	scoredMoves := ScoredMoves{}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"b179006e-e165-43a5-ac39-614dd4ec9a7e","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":278,"board":{"height":11,"width":11,"food":[{"x":6,"y":10},{"x":10,"y":4},{"x":10,"y":1},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":5},{"x":6,"y":0},{"x":0,"y":0},{"x":9,"y":0},{"x":10,"y":6},{"x":10,"y":10},{"x":0,"y":9},{"x":3,"y":0},{"x":0,"y":10},{"x":7,"y":0},{"x":8,"y":0},{"x":0,"y":4},{"x":10,"y":2},{"x":10,"y":3},{"x":7,"y":6},{"x":10,"y":0}],"hazards":[],"snakes":[{"id":"fcb38f74-3921-4982-a760-c948d2185a27","name":"Go Starter Project","health":96,"body":[{"x":5,"y":9},{"x":6,"y":9},{"x":6,"y":8},{"x":6,"y":7},{"x":6,"y":6},{"x":6,"y":5},{"x":6,"y":4},{"x":6,"y":3},{"x":6,"y":2},{"x":7,"y":2},{"x":8,"y":2},{"x":8,"y":3},{"x":7,"y":3},{"x":7,"y":4},{"x":8,"y":4},{"x":8,"y":5},{"x":9,"y":5},{"x":9,"y":6}],"head":{"x":5,"y":9},"length":18,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"8346a240-506b-4117-a3e9-1e114bc287c5","name":"ts-snake","health":100,"body":[{"x":5,"y":3},{"x":5,"y":4},{"x":5,"y":5},{"x":5,"y":6},{"x":5,"y":7},{"x":5,"y":8},{"x":4,"y":8},{"x":4,"y":7},{"x":4,"y":6},{"x":3,"y":6},{"x":3,"y":5},{"x":3,"y":4},{"x":3,"y":3},{"x":3,"y":2},{"x":3,"y":1},{"x":2,"y":1},{"x":2,"y":1}],"head":{"x":5,"y":3},"length":17,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"fcb38f74-3921-4982-a760-c948d2185a27","name":"Go Starter Project","health":96,"body":[{"x":5,"y":9},{"x":6,"y":9},{"x":6,"y":8},{"x":6,"y":7},{"x":6,"y":6},{"x":6,"y":5},{"x":6,"y":4},{"x":6,"y":3},{"x":6,"y":2},{"x":7,"y":2},{"x":8,"y":2},{"x":8,"y":3},{"x":7,"y":3},{"x":7,"y":4},{"x":8,"y":4},{"x":8,"y":5},{"x":9,"y":5},{"x":9,"y":6}],"head":{"x":5,"y":9},"length":18,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)
	//fmt.Printf("%v", state)
	//fmt.Printf("your head: %v", state.You.Head)
	//fmt.Printf("your neck: %v", state.You.Body[1])
	scoredMoves = AvoidSnakes(state, scoredMoves)
	expected := -1000
	actual := scoredMoves.Down.Score

	if actual != expected {
		t.Errorf("Expected: %v, got %v", expected, actual)

	}
}

func TestAvoidSnakeLeft(t *testing.T) {
	scoredMoves := ScoredMoves{}

	state := GameState{}
	json.Unmarshal([]byte(`{"game":{"id":"b179006e-e165-43a5-ac39-614dd4ec9a7e","ruleset":{"name":"standard","version":"cli","settings":{"foodSpawnChance":15,"minimumFood":1,"hazardDamagePerTurn":14}},"map":"standard","source":"","timeout":500},"turn":275,"board":{"height":11,"width":11,"food":[{"x":6,"y":10},{"x":10,"y":4},{"x":10,"y":1},{"x":2,"y":10},{"x":1,"y":10},{"x":0,"y":5},{"x":6,"y":0},{"x":0,"y":0},{"x":9,"y":0},{"x":10,"y":6},{"x":10,"y":10},{"x":0,"y":9},{"x":3,"y":0},{"x":0,"y":10},{"x":7,"y":0},{"x":8,"y":0},{"x":0,"y":4},{"x":10,"y":2},{"x":5,"y":3},{"x":10,"y":3},{"x":7,"y":6},{"x":10,"y":0}],"hazards":[],"snakes":[{"id":"fcb38f74-3921-4982-a760-c948d2185a27","name":"Go Starter Project","health":99,"body":[{"x":6,"y":7},{"x":6,"y":6},{"x":6,"y":5},{"x":6,"y":4},{"x":6,"y":3},{"x":6,"y":2},{"x":7,"y":2},{"x":8,"y":2},{"x":8,"y":3},{"x":7,"y":3},{"x":7,"y":4},{"x":8,"y":4},{"x":8,"y":5},{"x":9,"y":5},{"x":9,"y":6},{"x":9,"y":7},{"x":9,"y":8},{"x":9,"y":9}],"head":{"x":6,"y":7},"length":18,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}},{"id":"8346a240-506b-4117-a3e9-1e114bc287c5","name":"ts-snake","health":87,"body":[{"x":5,"y":6},{"x":5,"y":7},{"x":5,"y":8},{"x":4,"y":8},{"x":4,"y":7},{"x":4,"y":6},{"x":3,"y":6},{"x":3,"y":5},{"x":3,"y":4},{"x":3,"y":3},{"x":3,"y":2},{"x":3,"y":1},{"x":2,"y":1},{"x":1,"y":1},{"x":1,"y":2},{"x":1,"y":3}],"head":{"x":5,"y":6},"length":16,"latency":"0","shout":"","customizations":{"color":"#888888","head":"default","tail":"default"}}]},"you":{"id":"fcb38f74-3921-4982-a760-c948d2185a27","name":"Go Starter Project","health":99,"body":[{"x":6,"y":7},{"x":6,"y":6},{"x":6,"y":5},{"x":6,"y":4},{"x":6,"y":3},{"x":6,"y":2},{"x":7,"y":2},{"x":8,"y":2},{"x":8,"y":3},{"x":7,"y":3},{"x":7,"y":4},{"x":8,"y":4},{"x":8,"y":5},{"x":9,"y":5},{"x":9,"y":6},{"x":9,"y":7},{"x":9,"y":8},{"x":9,"y":9}],"head":{"x":6,"y":7},"length":18,"latency":"0","shout":"","customizations":{"color":"#2F2F2F","head":"default","tail":"default"}}}`), &state)
	//fmt.Printf("%v", state)
	//fmt.Printf("your head: %v", state.You.Head)
	//fmt.Printf("your neck: %v", state.You.Body[1])
	scoredMoves = AvoidSnakes(state, scoredMoves)
	expected := -1000
	actual := scoredMoves.Left.Score

	if actual != expected {
		t.Errorf("Expected: %v, got %v", expected, actual)

	}
}

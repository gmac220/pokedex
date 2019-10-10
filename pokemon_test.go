package main

import "testing"

func TestNewPokemon(t *testing.T) {
	testPokemon := NewPokemon(1, "Bulbasaur", true)

	if testPokemon.name == "Bulbasaur" {
		t.Log("the test Pokemon Bulbasaur was properly named")
	} else {
		t.Errorf("the test Pokemon was not Bulbasaur, but %s", testPokemon.name)
	}
}

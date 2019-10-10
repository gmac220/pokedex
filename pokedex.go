package main

import "fmt"

// Pokedex is directory of Pokemon
type Pokedex struct {
	directory []Pokemon
}

// Add will insert a new Pokemon into the drectory
func (p *Pokedex) Add(pokemon Pokemon) {
	p.directory = append(p.directory, pokemon)
}

// List will print out all Pokemon in the directory
func (p *Pokedex) List() {
	for _, pokemon := range p.directory {
		fmt.Println(pokemon.name, pokemon.caught)
	}
}

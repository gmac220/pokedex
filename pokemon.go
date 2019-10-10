package main

// Pokemon is a data structure representing each
// Pokemon we have seen or caught
type Pokemon struct {
	id     int
	name   string
	caught bool
}

// NewPokemon is a constructor for a Pokemon struct
func NewPokemon(id int, name string, caught bool) Pokemon {
	return Pokemon{
		id:     id,
		name:   name,
		caught: caught,
	}
}

package main

//var pokedex []Pokemon

func main() {
	pokedex := Pokedex{}
	bulbasaur := NewPokemon(1, "Bulbasaur", true)
	pokedex.Add(bulbasaur)
	pokedex.List()

}

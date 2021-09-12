package pokemon

type PokemonFiled struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Species string `json:"species"`
}

type PokemonInterface interface {
	GetPokemons() ([]PokemonFiled, error)
	GetPokemonById(id int) (*PokemonFiled, error)
	AddPokemon(data *PokemonFiled) (*PokemonFiled, error)
}

type PokemonController interface {
	GetPokemons() ([]PokemonFiled, error)
	GetPokemonById(id int) (*PokemonFiled, error)
	AddPokemon(data *PokemonFiled) (*PokemonFiled, error)
}
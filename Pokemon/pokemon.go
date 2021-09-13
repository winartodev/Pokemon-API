package pokemon

type Entity struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Species string `json:"species"`
}

type ModelInterface interface {
	GetPokemons() ([]Entity, error)
	GetPokemonByID(id int) (*Entity, error)
	AddPokemon(data *Entity) (*Entity, error)
}

type ControllerInterface interface {
	GetPokemons() ([]Entity, error)
	GetPokemonByID(id int) (*Entity, error)
	AddPokemon(data *Entity) (*Entity, error)
}
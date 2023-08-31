package entities

type PokemonList struct {
	Count   int           `json:"count"`
	Results []PokemonItem `json:"results"`
}

type PokemonItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonDetail struct {
	Base_experience int    `json:"base_experience"`
	Height          int    `json:"height"`
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Order           int    `json:"order"`
	Weight          int    `json:"weight"`
}

type Req struct {
	Url string `json:"url"`
}

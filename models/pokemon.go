package models

type Ability struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Abilities struct {
	Ability  Ability `json:"ability"`
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
}

type Pokemon struct {
	Id             int         `json:"id"`
	Name           string      `json:"name"`
	Abilities      []Abilities `json:"abilities"`
	Height         int         `json:"height"`
	Weight         int         `json:"weight"`
	BaseExperience int         `json:"base_experience"`
	Url            string      `json:"url"`
}

type Pokemons struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Pokemon `json:"results"`
}

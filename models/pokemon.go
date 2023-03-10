package models

type Ability struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}
type Abilities struct {
	Ability  Ability `json:"ability"`
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
}

type Type struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}

type Stat struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Pokemon struct {
	Id             int         `json:"id"`
	Name           string      `json:"name"`
	Abilities      []Abilities `json:"abilities"`
	Height         int         `json:"height"`
	Weight         int         `json:"weight"`
	Types          []Types     `json:"types"`
	Stats          []Stats     `json:"stats"`
	BaseExperience int         `json:"base_experience"`
	Url            string      `json:"url"`
}

type Pokemons struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Pokemon `json:"results"`
}

package controllers

type PokemonData struct {
	Id             int      `json:"id"`
	Name           string   `json:"name"`
	Abilities      []string `json:"abilities"`
	Height         int      `json:"height"`
	Weight         int      `json:"weight"`
	BaseExperience int      `json:"base_experience"`
}

type CompetitionData struct {
	Id       int `json:"id"`
	Rank1st  int `json:"rank1st"`
	Rank2nd  int `json:"rank2nd"`
	Rank3rd  int `json:"rank3rd"`
	Rank4th  int `json:"rank4th"`
	Rank5th  int `json:"rank5th"`
	SeasonId int `json:"season_id"`
}

type Season struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

type Pokemon struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type DataCompetition struct {
	Id      int         `json:"id"`
	Rank1st Pokemon     `json:"rank1st"`
	Rank2nd Pokemon     `json:"rank2nd"`
	Rank3rd Pokemon     `json:"rank3rd"`
	Rank4th Pokemon     `json:"rank4th"`
	Rank5th Pokemon     `json:"rank5th"`
	Season  interface{} `json:"season,omitempty"`
}

type DataScores struct {
	Id           int         `json:"id,omitempty"`
	Pokemon      Pokemon     `json:"pokemon"`
	Rank1stCount int         `json:"rank1st_count"`
	Rank2ndCount int         `json:"rank2nd_count"`
	Rank3rdCount int         `json:"rank3rd_count"`
	Rank4thCount int         `json:"rank4th_count"`
	Rank5thCount int         `json:"rank5th_count"`
	TotalPoint   int         `json:"total_point"`
	Season       interface{} `json:"season,omitempty"`
}

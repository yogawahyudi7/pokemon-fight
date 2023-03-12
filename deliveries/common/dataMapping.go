package common

type Stats struct {
	Name     string `json:"name"`
	BaseStat int    `json:"base_stat"`
	Effort   int    `json:"effort"`
}
type PokemonData struct {
	Id             int      `json:"id"`
	Name           string   `json:"name"`
	Abilities      []string `json:"abilities"`
	Height         int      `json:"height"`
	Weight         int      `json:"weight"`
	Types          []string `json:"types"`
	Stats          []Stats  `json:"stats"`
	BaseExperience int      `json:"base_experience"`
}

type CompetitionData struct {
	Id       int `json:"id"`
	Rank1st  int `json:"rank_1st"`
	Rank2nd  int `json:"rank_2nd"`
	Rank3rd  int `json:"rank_3rd"`
	Rank4th  int `json:"rank_4th"`
	Rank5th  int `json:"rank_5th"`
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
	Rank1st Pokemon     `json:"rank_1st"`
	Rank2nd Pokemon     `json:"rank_2nd"`
	Rank3rd Pokemon     `json:"rank_3rd"`
	Rank4th Pokemon     `json:"rank_4th"`
	Rank5th Pokemon     `json:"rank_5th"`
	Season  interface{} `json:"season,omitempty"`
}

type DataScores struct {
	Id           int         `json:"id,omitempty"`
	TotalPoints  int         `json:"total_points"`
	Pokemon      Pokemon     `json:"pokemon"`
	Rank1stCount int         `json:"rank_1st_count"`
	Rank2ndCount int         `json:"rank_2nd_count"`
	Rank3rdCount int         `json:"rank_3rd_count"`
	Rank4thCount int         `json:"rank_4th_count"`
	Rank5thCount int         `json:"rank_5th_count"`
	Season       interface{} `json:"season,omitempty"`
}

type UserOutput struct {
	ID    uint   `json:"id"`
	Level string `json:"level"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token,omitempty"`
}

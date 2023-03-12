package models

type Level struct {
	ID   uint   `gorm:"primary_key"`
	Name string `json:"name" form:"name"`
}

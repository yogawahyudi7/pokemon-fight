package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	LevelId  int    `json:"level_id" form:"level_id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Token    string `json:"token" form:"token"`

	DataLevel Level `gorm:"foreignKey:LevelId"`

	//FK
}

func (User) TableName() string {
	return "user"
}

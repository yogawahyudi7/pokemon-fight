package models

type Level struct {
	ID   uint   `gorm:"primary_key"`
	Name string `json:"name" form:"name"`

	//1 to many
	User []User `gorm:"foreignKey:LevelID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
}

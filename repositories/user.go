package repositories

import (
	"pokemon-fight/deliveries/middleware"
	"pokemon-fight/models"

	"gorm.io/gorm"
)

type UserRepositoriesInterface interface {
	CheckEmail(email string) (bool, error)
	GetLevel(id int) (models.Level, error)
	GetPassword(email string) (string, error)
	GetUserById(id int) (models.User, error)
	Register(user models.User) (models.User, error)
	Login(email string) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
}

type UserRepositories struct {
	db *gorm.DB
}

func NewUserRepositories(db *gorm.DB) *UserRepositories {
	return &UserRepositories{
		db: db,
	}
}

func (pr *UserRepositories) CheckEmail(email string) (bool, error) {
	var user models.User

	if err := pr.db.Model(&user).Where("email = ?", email).First(&user).Error; err != nil {
		return false, err
	}

	if user.Email == email {
		return true, nil
	} else {
		return false, nil
	}
}

func (pr *UserRepositories) GetLevel(id int) (models.Level, error) {
	var level models.Level
	if err := pr.db.Where("id = ?", id).First(&level).Error; err != nil {
		return level, err
	}

	return level, nil
}

func (pr *UserRepositories) GetPassword(email string) (string, error) {
	var user models.User

	query := pr.db.Debug()

	query = query.Where("email = ?", email)

	err := query.First(&user).Error

	if err != nil {
		return user.Password, err
	}
	return user.Password, nil
}

func (pr *UserRepositories) GetUserById(id int) (data models.User, err error) {

	if err = pr.db.Where("id = ?", id).First(&data).Error; err != nil {
		return data, err
	}

	return data, err
}

func (pr *UserRepositories) Register(user models.User) (models.User, error) {
	if err := pr.db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (pr *UserRepositories) Login(email string) (models.User, error) {
	var user models.User
	var err error
	if err = pr.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	user.Token, err = middleware.CreateToken(int(user.ID))
	if err != nil {
		return user, err
	}
	if err := pr.db.Save(user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (pr *UserRepositories) UpdateUser(user models.User) (models.User, error) {
	if err := pr.db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

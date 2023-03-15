package controllers

import (
	"fmt"
	"net/http"
	"pokemon-fight/deliveries/common"
	"pokemon-fight/deliveries/middleware"
	"pokemon-fight/helpers"
	"pokemon-fight/models"
	"pokemon-fight/repositories"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserControllers struct {
	Repositories repositories.UserRepositoriesInterface
}

func NewUserControllers(repositories repositories.UserRepositoriesInterface) *UserControllers {
	return &UserControllers{Repositories: repositories}
}

func (pc UserControllers) RegisterBos(ctx echo.Context) error {
	response := common.Response{}

	//get user's input
	userInput := models.User{}
	userInput.LevelId = 1
	ctx.Bind(&userInput)

	//check is data nil?
	if userInput.Email == "" || userInput.Password == "" || userInput.Name == "" || userInput.LevelId == 0 {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Dimohon Untuk Melengkapi Semua Data."))
	}

	err := validate.Var(userInput.Email, "email")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Parameter [email] Tidak Sesuai Format. Ex:email@xxx.com"))
	}

	//check is email exists?
	is_email_exists, _ := pc.Repositories.CheckEmail(userInput.Email)
	if is_email_exists {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Email Sudah Pernah Terdaftar."))
	}

	//encrypt pass user
	convert_pwd := []byte(userInput.Password) //convert pass from string to byte
	hashed_pwd := helpers.EncryptPwd(convert_pwd)
	userInput.Password = hashed_pwd //set new pass

	//create new user
	user, err := pc.Repositories.Register(userInput)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}
	fmt.Println("===========================", user.LevelId)
	//get level name
	level, err := pc.Repositories.GetLevel(int(user.LevelId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	//customize output
	result := common.UserOutput{
		ID:    user.ID,
		Level: level.Name,
		Email: user.Email,
		Name:  user.Name,
	}

	return ctx.JSON(http.StatusOK, response.Saved(result))
}

func (pc UserControllers) RegisterOperasional(ctx echo.Context) error {
	response := common.Response{}

	//get user's input
	userInput := models.User{}
	userInput.LevelId = 2
	ctx.Bind(&userInput)

	//check is data nil?
	if userInput.Email == "" || userInput.Password == "" || userInput.Name == "" || userInput.LevelId == 0 {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Dimohon Untuk Melengkapi Semua Data."))
	}

	err := validate.Var(userInput.Email, "email")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Parameter [email] Tidak Sesuai Format. Ex:email@xxx.com"))
	}

	//check is email exists?
	is_email_exists, _ := pc.Repositories.CheckEmail(userInput.Email)
	if is_email_exists {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Email Sudah Pernah Terdaftar."))
	}

	//encrypt pass user
	convert_pwd := []byte(userInput.Password) //convert pass from string to byte
	hashed_pwd := helpers.EncryptPwd(convert_pwd)
	userInput.Password = hashed_pwd //set new pass

	//create new user
	user, err := pc.Repositories.Register(userInput)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	//get level name
	level, err := pc.Repositories.GetLevel(int(user.LevelId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	//customize output
	result := common.UserOutput{
		ID:    user.ID,
		Level: level.Name,
		Email: user.Email,
		Name:  user.Name,
	}

	return ctx.JSON(http.StatusOK, response.Saved(result))
}

func (pc UserControllers) RegisterPengedar(ctx echo.Context) error {
	response := common.Response{}

	//get user's input
	userInput := models.User{}
	userInput.LevelId = 3
	ctx.Bind(&userInput)

	//check is data nil?
	if userInput.Email == "" || userInput.Password == "" || userInput.Name == "" || userInput.LevelId == 0 {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Dimohon Untuk Melengkapi Semua Data."))
	}

	err := validate.Var(userInput.Email, "email")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Parameter [email] Tidak Sesuai Format. Ex:email@xxx.com"))
	}

	//check is email exists?
	is_email_exists, _ := pc.Repositories.CheckEmail(userInput.Email)
	if is_email_exists {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Email Sudah Pernah Terdaftar."))
	}

	//encrypt pass user
	convert_pwd := []byte(userInput.Password) //convert pass from string to byte
	hashed_pwd := helpers.EncryptPwd(convert_pwd)
	userInput.Password = hashed_pwd //set new pass

	//create new user
	user, err := pc.Repositories.Register(userInput)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	//get level name
	level, err := pc.Repositories.GetLevel(int(user.LevelId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	//customize output
	result := common.UserOutput{
		ID:    user.ID,
		Level: level.Name,
		Email: user.Email,
		Name:  user.Name,
	}

	return ctx.JSON(http.StatusOK, response.Saved(result))
}

func (pc UserControllers) Login(ctx echo.Context) error {
	response := common.Response{}

	//get user's input
	userInput := models.User{}
	ctx.Bind(&userInput)

	//check is data nil?
	if userInput.Email == "" || userInput.Password == "" {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Dimohon Untuk Melengkapi Semua Data."))
	}

	err := validate.Var(userInput.Email, "email")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Parameter [email] Tidak Sesuai Format. Ex:email@xxx.com"))
	}

	//compare password on form with db
	get_pwd, err := pc.Repositories.GetPassword(userInput.Email) //get password
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(4))
		}
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}
	err = bcrypt.CompareHashAndPassword([]byte(get_pwd), []byte(userInput.Password))
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(4))
	}

	//login
	user, err := pc.Repositories.Login(userInput.Email)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(4))
		}
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	//get level name
	level, err := pc.Repositories.GetLevel(int(user.LevelId))
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(5))
		}
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	//customize output
	result := common.UserOutput{
		ID:    user.ID,
		Level: level.Name,
		Email: user.Email,
		Name:  user.Name,
		Token: user.Token,
	}

	return ctx.JSON(http.StatusOK, response.Login(result))
}

func (pc UserControllers) Logout(ctx echo.Context) error {
	response := common.Response{}

	//check otorisasi
	tokenUserId := middleware.ExtractToken(ctx)
	if tokenUserId == 0 {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(1))
	}
	user, err := pc.Repositories.GetUserById(tokenUserId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(6))
		}
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError("User Id : "+err.Error()))
	}

	user.Token = ""
	ctx.Bind(&user)
	customer_updated, err := pc.Repositories.UpdateUser(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	//get level name
	level, err := pc.Repositories.GetLevel(int(customer_updated.LevelId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	//customize output
	result := common.UserOutput{
		ID:    user.ID,
		Level: level.Name,
		Email: user.Email,
		Name:  user.Name,
		Token: user.Token,
	}

	return ctx.JSON(http.StatusOK, response.Logout(result))
}

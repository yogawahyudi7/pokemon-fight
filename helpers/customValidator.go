package helpers

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type ValidateDateStruct struct {
	DateString string `validate:"dateString"`
}

func ValidateDate(fl validator.FieldLevel) bool {
	regexExp := regexp.MustCompile(`^[0-9][0-9][0-9][0-9][-][0-9][0-9][-][0-9][0-9]$`)
	checking := regexExp.MatchString(fl.Field().String())

	fmt.Println("String Date : ", fl.Field().String())
	fmt.Println("Result Date : ", checking)

	return checking
}

func ValidatorDate(params string) (code int, err error) {

	validate = validator.New()

	if params == "" {
		return 1, errors.New("parameter tanggal tidak boleh kosong")
	}

	//helper validator
	validate.RegisterValidation("dateString", ValidateDate)
	str := ValidateDateStruct{DateString: params}

	//validasi numeric
	errsValid := validate.Struct(str)
	if errsValid != nil {

		fmt.Println("ERR : ", errsValid.Error())

		return 2, errors.New("format parameter tanggal tidak sesuai. ex:yyyy-mm-dd")

	}

	return 0, err
}

func ValidatorGeneralName(params string) (code int, err error) {

	validate = validator.New()

	if params == "" {
		return 1, errors.New("parameter nama tidak boleh kosong")
	}

	return 0, err
}

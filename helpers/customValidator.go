package helpers

import (
	"fmt"
	"pokemon-fight/constants"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/now"
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

func ValidatorDate(params string) (code int, err string) {

	validate = validator.New()

	if params == "" {
		return 1, ("Maaf, Parameter Tanggal Tidak Boleh Kosong.")
	}

	//helper validator
	validate.RegisterValidation("dateString", ValidateDate)
	str := ValidateDateStruct{DateString: params}

	//validasi numeric
	errsValid := validate.Struct(str)
	if errsValid != nil {

		fmt.Println("ERR : ", errsValid.Error())

		return 2, ("Maaf, Format Parameter Tanggal Tidak Sesuai. Ex:yyyy-mm-dd.")

	}

	splitDate := strings.Split(params, "-")
	year := splitDate[0]
	month := splitDate[1]
	day := splitDate[2]

	intYear, _ := strconv.Atoi(year)

	switch {
	case intYear < 1753:
		return 3, ("Maaf, Tahun Pada Parameter Tanggal Tidak Boleh Kecil Dari 1753.")
	}

	intMonth, _ := strconv.Atoi(month)
	switch {
	case intMonth > 12 || intMonth < 1:
		return 4, ("Maaf, Bulan Pada Parameter Tanggal Tidak Sesuai Dengan Bulan Yang Tersedia.")
	}

	layoutFormat := constants.LayoutYMD
	date, _ := time.Parse(layoutFormat, params)
	lastDate := fmt.Sprint(now.With(date).EndOfMonth().Format(layoutFormat))
	splitLastDate := strings.Split(lastDate, "-")
	lastDay := splitLastDate[2]
	intLastDay, _ := strconv.Atoi(lastDay)

	intDay, _ := strconv.Atoi(day)
	switch {
	case intDay > intLastDay || intDay < 1:

		return 5, ("Maaf, Hari Pada Parameter Tanggal Tidak Sesuai Dengan Hari yang Tersedia.")
	}

	return 0, ""
}

func ValidatorGeneralName(params string) (code int, err string) {

	validate = validator.New()

	if params == "" {
		return 1, ("Maaf, Parameter Nama Tidak Boleh Kosong.")
	}

	return 0, ""
}

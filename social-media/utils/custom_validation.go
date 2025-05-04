package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func NotBlank(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	return strings.TrimSpace(str) != ""
}

package presenter

import (
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field string `json:"field"`
	Message   string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
			case "required":
					return "Must be filled"
			case "email":
					return "E-mail invalid"
	}
return "Erro desconhecido"
}
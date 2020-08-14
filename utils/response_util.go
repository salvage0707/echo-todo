package utils

import (
	"github.com/go-playground/validator/v10"
)

// ValidateErrorResponse バリデーションエラーレスポンス
type ValidateErrorResponse struct {
	Field string      `json:"field"`
	Tag   string      `json:"error_tag"`
	Value interface{} `json:"value"`
}

// ValidateResponse バリデーションエラーレスポンスの構築
func ValidateResponse(errs validator.ValidationErrors) (response []ValidateErrorResponse) {

	for _, ve := range errs {
		fe := ve.(validator.FieldError)
		ver := ValidateErrorResponse{}
		ver.Field = fe.Field()
		ver.Tag = fe.Tag()
		ver.Value = fe.Value()

		response = append(response, ver)
	}

	return
}

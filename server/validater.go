package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CustomValidator カスタムバリデータ
type CustomValidator struct {
	validator *validator.Validate
}

// Validate バリデーションチェック
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// SetupValidater validator設定
func SetupValidater(e *echo.Echo) {
	e.Validator = &CustomValidator{validator: validator.New()}
}

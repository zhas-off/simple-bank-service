package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/zhas-off/simple-bank-service/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}

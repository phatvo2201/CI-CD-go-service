package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/phatvo2201/simplebank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.AcceptedCurrency(currency)
	}
	return false
}

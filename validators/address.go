package validators

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"github.com/shopicano/shopicano-backend/errors"
)

type ReqAddressCreate struct {
	Name     string `json:"name" valid:"required"`
	Street   string `json:"street" valid:"required"`
	City     string `json:"city" valid:"required"`
	Country  string `json:"country" valid:"required"`
	Postcode string `json:"postcode" valid:"required"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func ValidateCreateAddress(ctx echo.Context) (*ReqAddressCreate, error) {
	pld := ReqAddressCreate{}
	if err := ctx.Bind(&pld); err != nil {
		return nil, err
	}

	ok, err := govalidator.ValidateStruct(&pld)
	if ok {
		return &pld, nil
	}

	ve := errors.ValidationError{}

	for k, v := range govalidator.ErrorsByField(err) {
		ve.Add(k, v)
	}

	return nil, &ve
}

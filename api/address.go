package api

import (
	"github.com/labstack/echo/v4"
	"github.com/shopicano/shopicano-backend/app"
	"github.com/shopicano/shopicano-backend/core"
	"github.com/shopicano/shopicano-backend/data"
	"github.com/shopicano/shopicano-backend/errors"
	"github.com/shopicano/shopicano-backend/middlewares"
	"github.com/shopicano/shopicano-backend/models"
	"github.com/shopicano/shopicano-backend/utils"
	"github.com/shopicano/shopicano-backend/validators"
	"net/http"
	"strconv"
	"time"
)

func RegisterAddressRoutes(g *echo.Group) {
	func(g echo.Group) {
		g.Use(middlewares.AuthUser)
		g.POST("/", createAddress)
		g.GET("/", listAddresses)
		g.GET("/:address_id/", getAddress)
		g.DELETE("/:address_id/", deleteAddress)
		g.PUT("/:address_id/", updateAddress)
	}(*g)
}

func createAddress(ctx echo.Context) error {
	userID := ctx.Get(utils.UserID).(string)

	req, err := validators.ValidateCreateAddress(ctx)

	resp := core.Response{}

	if err != nil {
		resp.Title = "Invalid data"
		resp.Status = http.StatusUnprocessableEntity
		resp.Code = errors.CollectionCreationDataInvalid
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	add := &models.Address{
		ID:        utils.NewUUID(),
		Name:      req.Name,
		Address:   req.Address,
		City:      req.City,
		Postcode:  req.Postcode,
		Phone:     req.Phone,
		Email:     req.Email,
		Country:   req.Country,
		UserID:    userID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	db := app.DB()

	au := data.NewAddressRepository()
	if err := au.CreateAddress(db, add); err != nil {
		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Code = errors.DatabaseQueryFailed
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	resp.Status = http.StatusCreated
	resp.Data = add
	return resp.ServerJSON(ctx)
}

func updateAddress(ctx echo.Context) error {
	userID := ctx.Get(utils.UserID).(string)
	addressID := ctx.Param("address_id")

	req, err := validators.ValidateCreateAddress(ctx)

	resp := core.Response{}

	if err != nil {
		resp.Title = "Invalid data"
		resp.Status = http.StatusUnprocessableEntity
		resp.Code = errors.CollectionCreationDataInvalid
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	add := &models.Address{
		ID:        addressID,
		Name:      req.Name,
		Address:   req.Address,
		City:      req.City,
		Postcode:  req.Postcode,
		Phone:     req.Phone,
		Email:     req.Email,
		Country:   req.Country,
		UserID:    userID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	db := app.DB()

	au := data.NewAddressRepository()
	if err := au.UpdateAddress(db, add); err != nil {
		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Code = errors.DatabaseQueryFailed
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	resp.Status = http.StatusOK
	resp.Data = add
	return resp.ServerJSON(ctx)
}

func getAddress(ctx echo.Context) error {
	userID := ctx.Get(utils.UserID).(string)
	addressID := ctx.Param("address_id")

	resp := core.Response{}

	db := app.DB()

	au := data.NewAddressRepository()
	a, err := au.GetAddress(db, userID, addressID)
	if err != nil {
		if errors.IsRecordNotFoundError(err) {
			resp.Title = "Address not found"
			resp.Status = http.StatusNotFound
			resp.Code = errors.AddressNotFound
			resp.Errors = err
			return resp.ServerJSON(ctx)
		}

		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Code = errors.DatabaseQueryFailed
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	resp.Status = http.StatusOK
	resp.Data = a
	return resp.ServerJSON(ctx)
}

func deleteAddress(ctx echo.Context) error {
	userID := ctx.Get(utils.UserID).(string)
	addressID := ctx.Param("address_id")

	resp := core.Response{}

	db := app.DB()

	au := data.NewAddressRepository()
	if err := au.DeleteAddress(db, userID, addressID); err != nil {
		if errors.IsRecordNotFoundError(err) {
			resp.Title = "Address not found"
			resp.Status = http.StatusNotFound
			resp.Code = errors.AddressNotFound
			resp.Errors = err
			return resp.ServerJSON(ctx)
		}

		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Code = errors.DatabaseQueryFailed
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	resp.Status = http.StatusNoContent
	return resp.ServerJSON(ctx)
}

func listAddresses(ctx echo.Context) error {
	userID := ctx.Get(utils.UserID).(string)
	pageQ := ctx.Request().URL.Query().Get("page")
	limitQ := ctx.Request().URL.Query().Get("limit")

	page, err := strconv.ParseInt(pageQ, 10, 64)
	if err != nil {
		page = 1
	}
	limit, err := strconv.ParseInt(limitQ, 10, 64)
	if err != nil {
		limit = 10
	}

	resp := core.Response{}

	db := app.DB()

	from := (page - 1) * limit
	au := data.NewAddressRepository()
	addresses, err := au.ListAddresses(db, userID, int(from), int(limit))
	if err != nil {
		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Code = errors.DatabaseQueryFailed
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	resp.Status = http.StatusOK
	resp.Data = addresses
	return resp.ServerJSON(ctx)
}

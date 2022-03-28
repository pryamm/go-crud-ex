package handler

import (
	"backend-c-payment-monitoring/exception"
	"backend-c-payment-monitoring/exception/validation"
	"backend-c-payment-monitoring/model"
	"backend-c-payment-monitoring/service"
	"backend-c-payment-monitoring/util"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllPayment(c *fiber.Ctx) error {
	limit, err := strconv.Atoi(c.FormValue("limit"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString("limit required."),
			Data:    nil,
		})
	}

	page, err := strconv.Atoi(c.FormValue("page"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString("page required."),
			Data:    nil,
		})
	}

	keyword := c.FormValue("keyword")

	set_paginate := model.Pagination{}
	set_paginate.Limit = limit
	set_paginate.Page = page
	set_paginate.Keyword = keyword
	set_paginate.Sort = "Id asc"

	responses, err := service.GetAllPayment(set_paginate)

	if err != nil {
		//error
		return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(model.Pagination{
		Code:       http.StatusOK,
		Message:    "Get Data Success",
		Error:      nil,
		Limit:      responses.Limit,
		Page:       responses.Page,
		Sort:       responses.Sort,
		TotalRows:  responses.TotalRows,
		TotalPages: responses.TotalPages,
		Keyword:    responses.Keyword,
		Data:       responses.Data,
	})
}

func GetPaymentById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	responses, err := service.GetPaymentById(id)

	if err != nil {
		//error
		return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}
	return c.Status(http.StatusOK).JSON(model.ApiResponse{
		Code:    http.StatusOK,
		Message: "Get Data Success",
		Error:   nil,
		Data:    model.FormatGetPaymentResponse(responses),
	})
}

func CreatePayment(c *fiber.Ctx) error {
	var input model.CreatePaymentRequest

	err := c.BodyParser(&input)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	//validate input
	validation.PaymentValidate(input)

	//claim token
	claims, err := util.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(http.StatusInternalServerError).JSON(model.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	unit_id := claims.UserId

	responses, err := service.CreatePayment(input, uint(unit_id))

	if err != nil {
		//error
		return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Create Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(model.ApiResponse{
		Code:    http.StatusCreated,
		Message: "Create Data Success",
		Error:   nil,
		Data:    model.FormatCreatePaymentResponse(responses),
	})
}

func UpdateStatusPayment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(model.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	var payload model.UpdateStatusPaymentRequest
	err = c.BodyParser(&payload)

	//validation
	validation.PaymentStatusValidate(payload)

	if err != nil {

		//error
		return c.Status(http.StatusNotFound).JSON(model.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	responses, err := service.UpdateStatusPayment(id, payload)

	if err != nil {
		//error
		return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Update Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(model.ApiResponse{
		Code:    http.StatusOK,
		Message: "Update Data Success",
		Error:   nil,
		Data:    model.FormatUpdatePaymentResponse(responses),
	})
}

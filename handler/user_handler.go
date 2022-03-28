package handler

import (
	"backend-c-payment-monitoring/exception"
	"backend-c-payment-monitoring/exception/validation"
	"backend-c-payment-monitoring/model"
	"backend-c-payment-monitoring/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
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

	responses, err := service.GetAllUser(set_paginate)

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

func GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	responses, err := service.GetUserById(id)

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
		Data:    model.FormatGetUserResponse(responses),
	})
}

func CreateUser(c *fiber.Ctx) error {
	var payload model.CreateUserRequest

	err := c.BodyParser(&payload)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	//validation
	validation.UserValidate(payload)

	responses, err := service.CreateUser(payload)

	if err != nil {
		//error
		if err.Error() == "unique" {
			return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
				Code:    http.StatusBadRequest,
				Message: "Something Wrong",
				Error:   exception.NewString("username: Username Available"),
				Data:    nil,
			})
		}
		return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Create Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(model.ApiResponse{
		Code:    http.StatusOK,
		Message: "Create Data Success",
		Error:   nil,
		Data:    model.FormatCreateUserResponse(responses),
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(model.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	var payload model.CreateUserRequest
	err = c.BodyParser(&payload)

	//validation
	validation.UserValidate(payload)

	if err != nil {
		//error
		return c.Status(http.StatusNotFound).JSON(model.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	responses, err := service.UpdateUser(id, payload)

	if err != nil {
		//error
		if err.Error() == "unique" {
			return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
				Code:    http.StatusBadRequest,
				Message: "Something Wrong",
				Error:   exception.NewString("username: Username Available"),
				Data:    nil,
			})
		}
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
		Data:    model.FormatCreateUserResponse(responses),
	})
}

func DeleteUser(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(model.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	responses := service.DeleteUser(id)

	if responses != true {
		//error
		return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Delete Data Failed",
			Error:   exception.NewString("Record Not Found"),
			Data:    false,
		})
	}

	return c.Status(http.StatusOK).JSON(model.ApiResponse{
		Code:    http.StatusOK,
		Message: "Delete Data Success",
		Error:   nil,
		Data:    responses,
	})
}

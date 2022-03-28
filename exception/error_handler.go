package exception

import (
	"backend-c-payment-monitoring/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NewString(s string) *string {
	return &s
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)

	if ok {
		return ctx.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			// Status: "BAD_REQUEST",
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   NewString(err.Error()),
			Data:    nil,
		})
	}

	return ctx.Status(http.StatusInternalServerError).JSON(model.ApiResponse{
		// Status: "INTERNAL_SERVER_ERROR",
		Code:    http.StatusInternalServerError,
		Message: "Something Wrong",
		Error:   NewString(err.Error()),
		Data:    nil,
	})
}

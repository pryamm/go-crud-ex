package middleware

import (
	"backend-c-payment-monitoring/exception"
	"backend-c-payment-monitoring/model"
	"backend-c-payment-monitoring/util"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

//Admin
func RolePermissionAdmin(c *fiber.Ctx) error {
	// Get claims from JWT.
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

	role_id := claims.Role
	if role_id != 1 {
		return c.Status(http.StatusForbidden).JSON(model.ApiResponse{
			Code:    http.StatusForbidden,
			Message: "Something Wrong",
			Error:   exception.NewString("Forbidden"),
			Data:    nil,
		})
	}

	return c.Next()
}

//Unit
func RolePermissionUnit(c *fiber.Ctx) error {
	// Get claims from JWT.
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

	role_id := claims.Role
	if role_id != 2 {
		return c.Status(http.StatusForbidden).JSON(model.ApiResponse{
			Code:    http.StatusForbidden,
			Message: "Something Wrong",
			Error:   exception.NewString("Forbidden"),
			Data:    nil,
		})
	}

	return c.Next()
}

//GS & AC
func RolePermissionGsAc(c *fiber.Ctx) error {
	// Get claims from JWT.
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

	role_id := claims.Role
	if role_id != 3 && role_id != 4 {
		return c.Status(http.StatusForbidden).JSON(model.ApiResponse{
			Code:    http.StatusForbidden,
			Message: "Something Wrong",
			Error:   exception.NewString("Forbidden"),
			Data:    nil,
		})
	}

	return c.Next()
}

package handler

import (
	"backend-c-payment-monitoring/exception"
	"backend-c-payment-monitoring/exception/validation"
	"backend-c-payment-monitoring/model"
	"backend-c-payment-monitoring/service"
	"backend-c-payment-monitoring/util"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(ctx *fiber.Ctx) error {
	var input model.LoginUserRequest
	err := ctx.BodyParser(&input)
	if err != nil {
		return err
	}

	//validation
	validation.LoginValidate(input)

	responses, err := service.Login(input)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Login Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	token, err := util.GenerateNewAccessToken(responses)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(model.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "Login Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	refreshToken, err := util.GenerateRefreshToken(responses)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(model.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "Login Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	err = service.CreateAuthentication(responses, refreshToken)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(model.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "Set Authentication Refresh Token Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	ctx.Append("refreshToken", refreshToken)

	return ctx.Status(http.StatusOK).JSON(model.ApiResponse{
		Code:    http.StatusOK,
		Message: "Login Success",
		Error:   nil,
		Data:    model.FormatLoginUserResponse(responses, token),
	})
}

func RefreshToken(ctx *fiber.Ctx) error {

	_, err := util.VerifyToken(ctx)
	if err != nil && err.Error() != "Token is expired" {
		return ctx.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Access Token Not Valid",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	refreshTokenOld, err := util.ExtractTokenMetadataRefresh(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Refresh Token Not Valid",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	userId := int(refreshTokenOld.UserId)

	responses, err := service.GetUserById(userId)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong on User Data",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	accessToken, err := util.GenerateNewAccessToken(responses)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(model.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "Create Access Token Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	refreshToken, err := util.GenerateRefreshToken(responses)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(model.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "Create Refresh Token Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	refreshTokenOldData := util.ExtractTokenRefresh(ctx)
	err = service.UpdateAuthentication(refreshTokenOldData, refreshToken)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(model.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "Set Authentication Refresh Token Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	ctx.Append("refreshToken", refreshToken)

	return ctx.Status(http.StatusOK).JSON(model.ApiResponse{
		Code:    http.StatusOK,
		Message: "Refresh Token Success",
		Error:   nil,
		Data:    model.FormatLoginUserResponse(responses, accessToken),
	})

}

func LogoutHandler(ctx *fiber.Ctx) error {
	refreshToken := util.ExtractTokenRefresh(ctx)

	err := service.Logout(refreshToken)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(model.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "Logout Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(model.ApiResponse{
		Code:    http.StatusOK,
		Message: "Logout Success",
		Error:   nil,
		Data:    nil,
	})
}

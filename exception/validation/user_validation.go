package validation

import (
	"backend-c-payment-monitoring/exception"
	"backend-c-payment-monitoring/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func UserValidate(request model.CreateUserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required.When(request.Name == "").Error("Name is Required")),
		validation.Field(&request.Username, validation.Required.When(request.Username == "").Error("Username is Required")),
		validation.Field(&request.Password, validation.Required.When(request.Password == "").Error("Password is Required")),
		validation.Field(&request.RoleID, validation.Required.When(request.RoleID == 0).Error("Role is Required")),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

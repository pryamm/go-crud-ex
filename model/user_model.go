package model

import "gorm.io/gorm"

//entity
type User struct {
	gorm.Model
	ID       uint   `gorm:"not null"`
	Name     string `gorm:"size:255;not null"`
	Username string `gorm:"size:255;not null"`
	Password string `gorm:"size:255;not null"`
	RoleID   uint   `gorm:"not null"`
	Role     Role
}

//request
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleID   uint   `json:"role" binding:"required"`
}

//response
type UserResponse struct {
	ID       uint         `json:"id"`
	Name     string       `json:"name"`
	Username string       `json:"username"`
	Role     RoleResponse `json:"role"`
}

type UserCreateResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func FormatGetAllUserResponse(users []User) []UserResponse {
	usersFormatter := []UserResponse{}

	for _, user := range users {
		userFormatter := UserResponse{}
		userFormatter.ID = user.ID
		userFormatter.Name = user.Name
		userFormatter.Username = user.Username
		userFormatter.Role.ID = user.Role.ID
		userFormatter.Role.Name = user.Role.Name

		usersFormatter = append(usersFormatter, userFormatter)
	}

	return usersFormatter
}

func FormatGetUserResponse(user User) UserResponse {

	userFormatter := UserResponse{}
	userFormatter.ID = user.ID
	userFormatter.Name = user.Name
	userFormatter.Username = user.Username
	userFormatter.Role.ID = user.Role.ID
	userFormatter.Role.Name = user.Role.Name

	return userFormatter
}

func FormatCreateUserResponse(user User) UserCreateResponse {

	userFormatter := UserCreateResponse{}
	userFormatter.ID = user.ID
	userFormatter.Name = user.Name
	userFormatter.Username = user.Username

	return userFormatter
}

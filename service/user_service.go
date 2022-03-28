package service

import (
	"backend-c-payment-monitoring/model"
	"backend-c-payment-monitoring/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func GetAllUser(pagination model.Pagination) (model.Pagination, error) {

	users, err := repository.FindAllUser(pagination)

	if err != nil {
		return users, err
	}

	return users, err
}

func GetUserById(id int) (model.User, error) {
	user, err := repository.FindUserById(id)

	if err != nil {
		return user, err
	}

	return user, err
}

func CreateUser(payload model.CreateUserRequest) (model.User, error) {
	//check username
	username, err := repository.FindUserByUsername(payload.Username)

	if err == nil {
		return username, errors.New("unique")
	}

	user := model.User{}
	user.Name = payload.Name
	user.Username = payload.Username

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.RoleID = payload.RoleID

	inserted, err := repository.CreateUser(user)

	if err != nil {
		return inserted, err
	}

	return inserted, nil
}

func UpdateUser(id int, payload model.CreateUserRequest) (model.User, error) {
	//check availability
	userid, err := repository.FindUserById(id)

	if err != nil {
		return userid, err
	}

	//checkuniquename
	userName, err := repository.FindUserByUsername(payload.Username)

	if err == nil {
		if userid.Username != payload.Username {
			return userName, errors.New("unique")
		}
	}

	user := model.User{}
	user.ID = userid.ID
	user.Name = payload.Name
	user.Username = payload.Username

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.RoleID = payload.RoleID

	updateUser, err := repository.UpdateUser(id, user)

	if err != nil {
		return updateUser, err
	}

	return updateUser, nil
}

func DeleteUser(id int) bool {
	//check availability
	_, err := repository.FindUserById(id)

	if err != nil {
		return false
	}

	_, err = repository.DeleteUser(id)

	if err != nil {
		return false
	}

	return true
}

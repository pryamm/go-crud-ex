package service

import (
	"backend-c-payment-monitoring/model"
	"backend-c-payment-monitoring/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Login(input model.LoginUserRequest) (model.User, error) {

	username := input.Username
	password := input.Password

	user, err := repository.FindUserByUsername(username)

	if err != nil {
		return user, errors.New("Username Not Found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("Username/Password was wrong")
	}

	return user, nil
}

func Logout(refreshToken string) error {

	authentication, err := repository.FindAuthenticationByRefreshToken(refreshToken)

	if err != nil {
		return errors.New("Token Not Found")
	}

	err = repository.DeleteAuthentication(authentication.RefreshToken)
	if err != nil {
		return errors.New("Delete Token Failed")
	}

	return nil
}

func CreateAuthentication(payload model.User, refreshToken string) error {

	authentication := model.Authentication{}
	authentication.UserID = payload.ID
	authentication.RefreshToken = refreshToken

	err := repository.CreateAuthentication(authentication)
	if err != nil {
		return err
	}

	return nil
}

func UpdateAuthentication(refreshTokenOld string, refreshTokenNew string) error {
	authenticationOld, err := repository.FindAuthenticationByRefreshToken(refreshTokenOld)
	if err != nil {
		return err
	}

	authentication := model.Authentication{}
	authentication.RefreshToken = refreshTokenNew

	err = repository.UpdateAuthentication(int(authenticationOld.ID), authentication)

	if err != nil {
		return err
	}

	return nil
}

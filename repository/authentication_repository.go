package repository

import (
	"backend-c-payment-monitoring/config"
	"backend-c-payment-monitoring/model"
)

func FindAuthenticationById(id int) (model.Authentication, error) {
	db := config.DB

	var authentication model.Authentication

	err := db.First(&authentication, id).Error
	if err != nil {
		return authentication, err
	}

	return authentication, nil
}

func FindAuthenticationByRefreshToken(refreshToken string) (model.Authentication, error) {
	db := config.DB

	var authentication model.Authentication

	err := db.Where("refresh_token = ?", refreshToken).First(&authentication).Error
	if err != nil {
		return authentication, err
	}

	return authentication, nil
}

func CreateAuthentication(authentication model.Authentication) error {
	db := config.DB

	err := db.Save(&authentication).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateAuthentication(id int, authentication model.Authentication) error {
	db := config.DB

	err := db.Where("id = ?", id).Updates(&authentication).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteAuthentication(refreshToken string) error {
	db := config.DB

	var authentication model.Authentication

	err := db.Where("refresh_token = ?", refreshToken).Delete(&authentication).Error
	if err != nil {
		return err
	}

	return nil
}

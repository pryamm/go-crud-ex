package repository

import (
	"backend-c-payment-monitoring/config"
	"backend-c-payment-monitoring/model"
)

func FindAllUser(pagination model.Pagination) (model.Pagination, error) {
	db := config.DB

	var users []model.User

	keyword := "%" + pagination.Keyword + "%"
	err := db.
		Preload("Role").
		Where("name LIKE ?", keyword).
		Scopes(model.UserPaginate(keyword, users, &pagination, db)).
		Find(&users).Error

	if err != nil {
		return pagination, err
	}

	pagination.Data = model.FormatGetAllUserResponse(users)

	return pagination, nil
}

func FindUserById(id int) (model.User, error) {
	db := config.DB

	var user model.User

	err := db.Preload("Role").First(&user, id).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func FindUserByUsername(username string) (model.User, error) {
	db := config.DB

	var user model.User

	err := db.Preload("Role").Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user model.User) (model.User, error) {
	db := config.DB

	err := db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUser(id int, user model.User) (model.User, error) {
	db := config.DB

	err := db.Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func DeleteUser(id int) (model.User, error) {
	db := config.DB

	var user model.User

	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

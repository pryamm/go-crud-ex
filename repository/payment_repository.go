package repository

import (
	"backend-c-payment-monitoring/config"
	"backend-c-payment-monitoring/model"
)

func FindAllPayment(pagination model.Pagination) (model.Pagination, error) {
	db := config.DB

	var payments []model.Payment

	keyword := "%" + pagination.Keyword + "%"
	err := db.
		Preload("Status").
		Where("request_by LIKE ?", keyword).
		Scopes(model.PaymentPaginate(keyword, payments, &pagination, db)).
		Find(&payments).Error

	if err != nil {
		return pagination, err
	}

	pagination.Data = model.FormatGetAllPaymentResponse(payments)

	return pagination, nil
}

func FindPaymentById(id int) (model.Payment, error) {
	db := config.DB

	var payment model.Payment

	err := db.Preload("User").Preload("Status").First(&payment, id).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func CreatePayment(payment model.Payment) (model.Payment, error) {
	db := config.DB

	err := db.Save(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func UpdateStatusPayment(id int, payment model.Payment) (model.Payment, error) {
	db := config.DB

	err := db.Model(&payment).Where("id = ?", id).Updates(map[string]interface{}{"status_id": payment.StatusID, "reason": payment.Reason}).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}

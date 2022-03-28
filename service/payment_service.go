package service

import (
	"backend-c-payment-monitoring/model"
	"backend-c-payment-monitoring/repository"
	"time"
)

func GetAllPayment(pagination model.Pagination) (model.Pagination, error) {

	users, err := repository.FindAllPayment(pagination)

	if err != nil {
		return users, err
	}

	return users, err
}

func GetPaymentById(id int) (model.Payment, error) {
	payment, err := repository.FindPaymentById(id)

	if err != nil {
		return payment, err
	}

	return payment, err
}

func CreatePayment(payload model.CreatePaymentRequest, unit_id uint) (model.Payment, error) {

	payment := model.Payment{}
	payment.UserID = unit_id
	payment.RequestBy = payload.RequestBy
	payment.Necessity = payload.Necessity

	t, _ := time.Parse("2006-01-02", payload.PaymentDate)

	payment.PaymentDate = t
	payment.PaymentAmount = payload.PaymentAmount
	payment.PaymentCalculate = payload.PaymentCalculate
	payment.PaymentAccountName = payload.PaymentAccountName
	payment.PaymentAccountNumber = payload.PaymentAccountNumber
	payment.StatusID = 1

	newPayment, err := repository.CreatePayment(payment)

	if err != nil {
		return newPayment, err
	}

	return newPayment, nil
}

func UpdateStatusPayment(id int, payload model.UpdateStatusPaymentRequest) (model.Payment, error) {
	//check id
	paymentid, err := repository.FindPaymentById(id)

	if err != nil {
		return paymentid, err
	}

	payment := model.Payment{}
	payment.ID = paymentid.ID
	payment.UserID = paymentid.UserID
	payment.RequestBy = paymentid.RequestBy
	payment.Necessity = paymentid.Necessity
	payment.PaymentDate = paymentid.PaymentDate
	payment.PaymentAmount = paymentid.PaymentAmount
	payment.PaymentCalculate = paymentid.PaymentCalculate
	payment.PaymentAccountName = paymentid.PaymentAccountName
	payment.PaymentAccountNumber = paymentid.PaymentAccountNumber
	payment.StatusID = payload.StatusID
	payment.Reason = payload.Reason

	updatePayment, err := repository.UpdateStatusPayment(id, payment)

	if err != nil {
		return updatePayment, err
	}

	return updatePayment, nil
}

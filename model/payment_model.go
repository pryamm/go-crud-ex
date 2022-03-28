package model

import (
	"time"

	"gorm.io/gorm"
)

//entity
type Payment struct {
	gorm.Model
	ID                   uint      `gorm:"not null"`
	UserID               uint      `gorm:"not null"`
	RequestBy            string    `gorm:"size:255;not null"`
	Necessity            string    `gorm:"size:255;not null"`
	PaymentDate          time.Time `gorm:"type:date;not null"`
	PaymentAmount        int       `gorm:"not null"`
	PaymentCalculate     string    `gorm:"not null"`
	PaymentAccountName   string    `gorm:"size:255;not null"`
	PaymentAccountNumber string    `gorm:"size:255;not null"`
	StatusID             uint      `gorm:"not null"`
	Reason               *string   `gorm:"default:null"`
	CreatedAt            time.Time
	User                 User
	Status               Status
}

//request
type CreatePaymentRequest struct {
	RequestBy            string `json:"request_by" binding:"required"`
	Necessity            string `json:"necessity" binding:"required"`
	PaymentDate          string `json:"payment_date" binding:"required"`
	PaymentAmount        int    `json:"payment_amount" binding:"required"`
	PaymentCalculate     string `json:"payment_calculate" binding:"required"`
	PaymentAccountName   string `json:"payment_account_name" binding:"required"`
	PaymentAccountNumber string `json:"payment_account_number" binding:"required"`
}

type UpdateStatusPaymentRequest struct {
	StatusID uint    `json:"status_id" binding:"required"`
	Reason   *string `json:"reason" binding:"required"`
}

//response
type PaymentResponse struct {
	ID                   uint           `json:"id"`
	UserID               uint           `json:"unit_id"`
	RequestBy            string         `json:"request_by"`
	Necessity            string         `json:"necessity"`
	PaymentDate          time.Time      `json:"payment_date"`
	PaymentAmount        int            `json:"payment_amount"`
	PaymentCalculate     string         `json:"payment_calculate"`
	PaymentAccountName   string         `json:"payment_account_name"`
	PaymentAccountNumber string         `json:"payment_account_number"`
	Reason               *string        `json:"reason"`
	CreatedAt            time.Time      `json:"created_at"`
	Status               StatusResponse `json:"status"`
}

type CreatePaymentResponse struct {
	ID                   uint      `json:"id"`
	UserID               uint      `json:"unit_id"`
	RequestBy            string    `json:"request_by"`
	Necessity            string    `json:"necessity"`
	PaymentDate          time.Time `json:"payment_date"`
	PaymentAmount        int       `json:"payment_amount"`
	PaymentCalculate     string    `json:"payment_calculate"`
	PaymentAccountName   string    `json:"payment_account_name"`
	PaymentAccountNumber string    `json:"payment_account_number"`
}

type UpdatePaymentResponse struct {
	ID                   uint      `json:"id"`
	UserID               uint      `json:"unit_id"`
	RequestBy            string    `json:"request_by"`
	Necessity            string    `json:"necessity"`
	PaymentDate          time.Time `json:"payment_date"`
	PaymentAmount        int       `json:"payment_amount"`
	PaymentCalculate     string    `json:"payment_calculate"`
	PaymentAccountName   string    `json:"payment_account_name"`
	PaymentAccountNumber string    `json:"payment_account_number"`
	Reason               *string   `json:"reason"`
}

func FormatGetAllPaymentResponse(payments []Payment) []PaymentResponse {
	paymentsFormatter := []PaymentResponse{}

	for _, payment := range payments {
		paymentFormatter := PaymentResponse{}
		paymentFormatter.ID = payment.ID
		paymentFormatter.UserID = payment.UserID
		paymentFormatter.RequestBy = payment.RequestBy
		paymentFormatter.Necessity = payment.Necessity
		paymentFormatter.PaymentDate = payment.PaymentDate
		paymentFormatter.PaymentAmount = payment.PaymentAmount
		paymentFormatter.PaymentCalculate = payment.PaymentCalculate
		paymentFormatter.PaymentAccountName = payment.PaymentAccountName
		paymentFormatter.PaymentAccountNumber = payment.PaymentAccountNumber
		paymentFormatter.Reason = payment.Reason
		paymentFormatter.Status.ID = payment.Status.ID
		paymentFormatter.Status.Name = payment.Status.Name
		paymentFormatter.CreatedAt = payment.CreatedAt

		paymentsFormatter = append(paymentsFormatter, paymentFormatter)
	}

	return paymentsFormatter
}

func FormatGetPaymentResponse(payment Payment) PaymentResponse {

	paymentFormatter := PaymentResponse{}
	paymentFormatter.ID = payment.ID
	paymentFormatter.UserID = payment.UserID
	paymentFormatter.RequestBy = payment.RequestBy
	paymentFormatter.Necessity = payment.Necessity
	paymentFormatter.PaymentDate = payment.PaymentDate
	paymentFormatter.PaymentAmount = payment.PaymentAmount
	paymentFormatter.PaymentCalculate = payment.PaymentCalculate
	paymentFormatter.PaymentAccountName = payment.PaymentAccountName
	paymentFormatter.PaymentAccountNumber = payment.PaymentAccountNumber
	paymentFormatter.Reason = payment.Reason
	paymentFormatter.Status.ID = payment.Status.ID
	paymentFormatter.Status.Name = payment.Status.Name
	paymentFormatter.CreatedAt = payment.CreatedAt

	return paymentFormatter
}

func FormatCreatePaymentResponse(payment Payment) CreatePaymentResponse {

	paymentFormatter := CreatePaymentResponse{}
	paymentFormatter.ID = payment.ID
	paymentFormatter.UserID = payment.UserID
	paymentFormatter.RequestBy = payment.RequestBy
	paymentFormatter.Necessity = payment.Necessity
	paymentFormatter.PaymentDate = payment.PaymentDate
	paymentFormatter.PaymentAmount = payment.PaymentAmount
	paymentFormatter.PaymentCalculate = payment.PaymentCalculate
	paymentFormatter.PaymentAccountName = payment.PaymentAccountName
	paymentFormatter.PaymentAccountNumber = payment.PaymentAccountNumber

	return paymentFormatter
}

func FormatUpdatePaymentResponse(payment Payment) UpdatePaymentResponse {

	paymentFormatter := UpdatePaymentResponse{}
	paymentFormatter.ID = payment.ID
	paymentFormatter.UserID = payment.UserID
	paymentFormatter.RequestBy = payment.RequestBy
	paymentFormatter.Necessity = payment.Necessity
	paymentFormatter.PaymentDate = payment.PaymentDate
	paymentFormatter.PaymentAmount = payment.PaymentAmount
	paymentFormatter.PaymentCalculate = payment.PaymentCalculate
	paymentFormatter.PaymentAccountName = payment.PaymentAccountName
	paymentFormatter.PaymentAccountNumber = payment.PaymentAccountNumber
	paymentFormatter.Reason = payment.Reason

	return paymentFormatter
}

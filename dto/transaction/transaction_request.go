package transactiondto

import "time"

type TransactionRequest struct {
	StartDate time.Time `json:"startDate"`
	DueDate   time.Time `json:"dueDate"`
	UserID    int       `json:"userId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Price     int       `json:"price" validate:"required"`
}

type UpdateTransactionRequest struct {
	StartDate time.Time `json:"startDate" `
	DueDate   time.Time `json:"dueDate" `
	UserID    int       `json:"userId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `
	Price     int       `json:"price"`
}

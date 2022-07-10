package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type FixedCost struct {
	gorm.Model
	Name        string    `json:"name" gorm:"column:name"`
	Type        string    `json:"type" gorm:"column:type"`
	DueDate     time.Time `json:"dueDate" gorm:"column:due_date"`
	Value       float64   `json:"value" gorm:"column:value"`
	PaymentType string    `json:"paymentType" gorm:"column:payment_type"`
	Account     string    `json:"account" gorm:"column:account"`
}

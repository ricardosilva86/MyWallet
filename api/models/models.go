package models

import (
	"github.com/jinzhu/gorm"
	"gorm.io/datatypes"
)

type FixedCost struct {
	gorm.Model
	Name        string         `json:"name" gorm:"column:name"`
	Type        string         `json:"type" gorm:"column:type"`
	DueDate     datatypes.Date `json:"dueDate" gorm:"column:due_date"`
	Value       float64        `json:"value" gorm:"column:value"`
	PaymentType string         `json:"paymentType" gorm:"column:payment_type"`
	Account     string         `json:"account" gorm:"column:account"`
}

type VariableSpend struct {
	gorm.Model
	Name        string         `json:"name" gorm:"column:name"`
	Type        string         `json:"type" gorm:"column:type"`
	Date        datatypes.Date `json:"date" gorm:"column:date"`
	Amount      float64        `json:"amount" gorm:"column:amount"`
	PaymentType string         `json:"paymentType" gorm:"column:payment_type"`
	Account     string         `json:"account" gorm:"column:account"`
}

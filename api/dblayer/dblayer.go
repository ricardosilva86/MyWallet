package dblayer

import (
	"MyWallet/api/models"
	"errors"
)

type DBLayer interface {
	GetAllFixedCosts() ([]models.FixedCost, error)
	GetAllVariableSpend() ([]models.VariableSpend, error)
}

var ErrINVALIDPASSWORD = errors.New("Invalid password")

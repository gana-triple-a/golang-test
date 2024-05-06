package domain

import (
	"github.com/shopspring/decimal"
)

type Account struct {
	AccountId      int             `json:"account_id"`
	InitialBalance decimal.Decimal `json:"initial_balance"`
}

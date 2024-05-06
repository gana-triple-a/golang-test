package domain

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type CreateAccountPayload struct {
	AccountId int
	InitialBalance decimal.Decimal
}

func (dmgr *Mgr) CreateAccount(cap *CreateAccountPayload) (*Account, error) {
	for _, acc := range dmgr.accounts {
		if acc.AccountId == cap.AccountId {
			return nil, fmt.Error("Account already exists")
		}
	}

	if (cap.InitialBalance.LessThanOrEqual(decimal.Zero)) {
		return nil, fmt.Error("Initial amount less than zero")
	}

	newAccount := Account{
		AccountId: cap.AccountId,
		InitialBalance: cap.InitialBalance,
	}

	dmgr.accountMu.Lock()
	append(dmgr.accounts, newAccount)
	dmgr.accountMu.Unlock()

	return newAccount, nil
}
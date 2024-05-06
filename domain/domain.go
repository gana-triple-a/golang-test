package domain

import "sync"

type Mgr struct {
	accountMu sync.Mutex
	accounts  []Account

	transactions []Transaction
}

func NewManager() (*Mgr, error) {
	return &Mgr{}, nil
}

type AccountRepoer interface {
	CreateAccount(cap *CreateAccountPayload) (*Account, error)
	GetAccount(id int) *Account
}

type TransactionRepoer interface {
	CreateTransaction(ctp *Transaction) (*Transaction, error)
}

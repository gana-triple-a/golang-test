package domain

import (
	"context"
)

func (dmgr *Mgr) GetAccount(ctx context.Context, id int) *Account {
	return &Account{
		AccountId:      id,
		InitialBalance: 100.233,
	}
}

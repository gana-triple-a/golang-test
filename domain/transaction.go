package domain

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Transaction {
	SourceAccountId      int
	DestinationAccountId int
	Amount               decimal.Decimal
}

func (mgr *Mgr) CreateTransaction(ctp *Transaction	// for _, acc := range mgr.accounts {
	// 	if acc.AccountId == ctp.SourceAccountId {
	// 		acc.InitialBalance -= ctp.Amount
	// 	}
	// 	if acc.AccountId == ctp.DestinationAccountId {
	// 		acc.InitialBalance += ctp.Amount
	// 	}
	// }
	// newTransaction := Transaction{
	// 	SourceAccountId
	// }		


	// append(mgr.transactions)

	fmt.Println("transaction created!")

}

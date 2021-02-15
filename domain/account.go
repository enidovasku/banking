package domain

import (
	"github.com/enidovasku/banking/dto"
	"github.com/enidovasku/banking/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindBy(string) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountId}
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount >= amount
}

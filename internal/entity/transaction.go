package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
	CreatedAt   time.Time
}

func NewTransaction(accountFrom *Account, accountTo *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		ID:          uuid.New().String(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}

	err := transaction.Validate()
	if err != nil {
		return nil, err
	}
	transaction.Execute()
	return transaction, nil
}

func (t *Transaction) Validate() error {
	if t.AccountFrom == nil {
		return errors.New("transaction account from is required")
	}
	if t.AccountTo == nil {
		return errors.New("transaction account to is required")
	}
	if t.Amount <= 0 {
		return errors.New("transaction amount must be greater than 0")
	}
	if t.AccountFrom.Balance < t.Amount {
		return errors.New("account from does not have enough balance")
	}
	return nil
}

func (t *Transaction) Execute() error {
	t.AccountFrom.Withdraw(t.Amount)
	t.AccountTo.Deposit(t.Amount)
	return nil
}

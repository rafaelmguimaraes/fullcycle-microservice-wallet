package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Client    *Client
	Balance   float64
	CreatedAt time.Time
	UpdateAt  time.Time
}

func NewAccount(client *Client) (*Account, error) {
	account := &Account{
		ID:        uuid.New().String(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	err := account.Validate()
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (a *Account) Validate() error {
	if a.Client == nil {
		return errors.New("account client is required")
	}
	if a.Balance < 0 {
		return errors.New("account balance must be greater than or equal to 0")
	}
	return nil
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than 0")
	}
	a.Balance += amount
	a.UpdateAt = time.Now()
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be greater than 0")
	}
	if a.Balance < amount {
		return errors.New("insufficient account balance")
	}
	a.Balance -= amount
	a.UpdateAt = time.Now()
	return nil
}

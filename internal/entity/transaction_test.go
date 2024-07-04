package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	clientJohn, _ := NewClient("John Doe", "jo@j.com")
	accountJohn, _ := NewAccount(clientJohn)
	clientJane, _ := NewClient("Jane Doe", "ja@j.com")
	accountJane, _ := NewAccount(clientJane)

	accountJohn.Deposit(100)
	accountJane.Deposit(100)

	transaction, err := NewTransaction(accountJohn, accountJane, 50)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, accountJohn.ID, transaction.AccountFrom.ID)
	assert.Equal(t, accountJane.ID, transaction.AccountTo.ID)
	assert.Equal(t, 50.0, transaction.Amount)
	assert.Equal(t, 50.0, accountJohn.Balance)
	assert.Equal(t, 150.0, accountJane.Balance)

}

func TestCreateTransactionWithInvalidAmount(t *testing.T) {
	clientJohn, _ := NewClient("John Doe", "jo@j.com")
	accountJohn, _ := NewAccount(clientJohn)
	clientJane, _ := NewClient("Jane Doe", "ja@j.com")
	accountJane, _ := NewAccount(clientJane)

	accountJohn.Deposit(100)
	accountJane.Deposit(100)

	transactionAmountInsuficient, err := NewTransaction(accountJohn, accountJane, 200)
	assert.NotNil(t, err)
	assert.Error(t, err, "account from does not have enough balance")
	assert.Nil(t, transactionAmountInsuficient)
	assert.Equal(t, 100.0, accountJohn.Balance)
	assert.Equal(t, 100.0, accountJane.Balance)

	transactionAmountNegative, err := NewTransaction(accountJohn, accountJane, -100)
	assert.NotNil(t, err)
	assert.Error(t, err, "transaction amount must be greater than 0")
	assert.Nil(t, transactionAmountNegative)
	assert.Equal(t, 100.0, accountJohn.Balance)
	assert.Equal(t, 100.0, accountJane.Balance)

}

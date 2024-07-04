package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account, err := NewAccount(client)
	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
	assert.Equal(t, client, account.Client)
	assert.Equal(t, 0.0, account.Balance)
}

func TestCreateNewAccountWithNilClient(t *testing.T) {
	account, err := NewAccount(nil)
	assert.NotNil(t, err)
	assert.Nil(t, account, "account client is required")
}

func TestDeposit(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account, _ := NewAccount(client)
	err := account.Deposit(100)
	assert.Nil(t, err)
	assert.Equal(t, 100.0, account.Balance)
}

func TestDepositWithNegativeAmount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account, _ := NewAccount(client)
	err := account.Deposit(-100)
	assert.NotNil(t, err, "deposit amount must be greater than 0")
}

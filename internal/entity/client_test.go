package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "j@j.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestCreateNewClientWithEmptyName(t *testing.T) {
	client, err := NewClient("", "j@j.com")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestCreateNewClientWithEmptyEmail(t *testing.T) {
	client, err := NewClient("John Doe", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("Jane Doe", "j@j.com")
	assert.Nil(t, err)
	assert.Equal(t, "Jane Doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestUpdateClientWithEmptyName(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("", "j@j.com")
	assert.NotNil(t, err, "client name is required")
}

func TestUpdateClientWithEmptyEmail(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("Jane Doe", "")
	assert.NotNil(t, err, "client email is required")
}

func TestAddAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account, _ := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
	assert.Equal(t, account, client.Accounts[0])
}

func TestAddAccountWithDifferentClient(t *testing.T) {
	clientJohn, _ := NewClient("John Doe", "j@j.com")
	clientJane, _ := NewClient("Jane Doe", "ja@j.com")
	account, _ := NewAccount(clientJane)
	err := clientJohn.AddAccount(account)
	assert.NotNil(t, err, "account does not belong to client")
	assert.Equal(t, 0, len(clientJohn.Accounts))
}

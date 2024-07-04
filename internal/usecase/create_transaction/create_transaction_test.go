package createtransaction

import (
	"testing"

	"github.com/rafaelmguimaraes/fullcycle-microservice-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Get(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	clientJohn, _ := entity.NewClient("John", "j@j.com")
	accountJohn, _ := entity.NewAccount(clientJohn)
	accountJohn.Deposit(1000)

	clientJane, _ := entity.NewClient("Jane", "ja@j.com")
	accountJane, _ := entity.NewAccount(clientJane)
	accountJane.Deposit(1000)

	mockAccountGateway := new(AccountGatewayMock)
	mockAccountGateway.On("Get", accountJohn.ID).Return(accountJohn, nil)
	mockAccountGateway.On("Get", accountJane.ID).Return(accountJane, nil)

	mockTransactionGateway := new(TransactionGatewayMock)
	mockTransactionGateway.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInputDTO{
		AccountIDFrom: accountJohn.ID,
		AccountIDTo:   accountJane.ID,
		Amount:        100,
	}

	uc := NewCreateTransactionUseCase(mockTransactionGateway, mockAccountGateway)

	output, err := uc.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	mockAccountGateway.AssertExpectations(t)
	mockTransactionGateway.AssertExpectations(t)
	mockAccountGateway.AssertNumberOfCalls(t, "Get", 2)
	mockTransactionGateway.AssertNumberOfCalls(t, "Create", 1)

}

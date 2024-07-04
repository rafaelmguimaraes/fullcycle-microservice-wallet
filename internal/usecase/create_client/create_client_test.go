package createclient

import (
	"testing"

	"github.com/rafaelmguimaraes/fullcycle-microservice-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func TestCreateClientUseCase_Execute(t *testing.T) {
	clientGatewayMock := new(ClientGatewayMock)
	clientGatewayMock.On("Save", mock.Anything).Return(nil)

	createClientUseCase := NewCreateClientUseCase(clientGatewayMock)

	input := CreateClientInputDTO{
		Name:  "Rafael",
		Email: "rafael@r.com",
	}

	output, err := createClientUseCase.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, input.Name, output.Name)
	assert.Equal(t, input.Email, output.Email)
	clientGatewayMock.AssertExpectations(t)
	clientGatewayMock.AssertNumberOfCalls(t, "Save", 1)

}

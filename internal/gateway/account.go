package gateway

import "github.com/rafaelmguimaraes/fullcycle-microservice-wallet/internal/entity"

type AccountGateway interface {
	Get(id string) (*entity.Account, error)
	Save(account *entity.Account) error
}

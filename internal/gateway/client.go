package gateway

import "github.com/rafaelmguimaraes/fullcycle-microservice-wallet/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}

package gateway

import "github.com/rafaelmguimaraes/fullcycle-microservice-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}

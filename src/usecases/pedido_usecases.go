package usecases

import (
	"github.com/gomesmatheus/tech-challenge/src/entities"
	"github.com/gomesmatheus/tech-challenge/src/interfaces"
)

type pedidoUseCases struct {
	gateway interfaces.PedidoGateway
}

func NewPedidoUseCases(pedidoGateway interfaces.PedidoGateway) *pedidoUseCases {
	return &pedidoUseCases{
		gateway: pedidoGateway,
	}
}

func (usecase *pedidoUseCases) CriarPedido(p entities.Pedido) (entities.Pedido, error) {
	return usecase.gateway.CriarPedido(p)
}
func (usecase *pedidoUseCases) RecuperarPedidos() ([]entities.Pedido, error) {
	return usecase.gateway.RecuperarPedidos()
}

func (usecase *pedidoUseCases) AtualizarStatus(id int, status string) error {
	return usecase.gateway.AtualizarStatus(id, status)
}

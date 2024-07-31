package usecases

import (

	"github.com/gomesmatheus/tech-challenge/src/interfaces"
)

type pagamentoUseCases struct {
   gateway interfaces.PedidoGateway
}

func NewPagamentoUseCases(pedidoGateway interfaces.PedidoGateway) *pagamentoUseCases{
    return &pagamentoUseCases{
        gateway: pedidoGateway,
    }
}

func (usecase *pagamentoUseCases) AtualizarPagamento(id int, status bool) (error) {
    return usecase.gateway.AtualizarPagamento(id, status)
}

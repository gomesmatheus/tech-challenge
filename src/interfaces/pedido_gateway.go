package interfaces

import "github.com/gomesmatheus/tech-challenge/src/entities"

type PedidoGateway interface {
	CriarPedido(entities.Pedido) (entities.Pedido, error)
	RecuperarPedidos() ([]entities.Pedido, error)
	AtualizarStatus(id int, status string) error
	AtualizarPagamento(id int, status bool) error
}

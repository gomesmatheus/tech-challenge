package interfaces

import "github.com/gomesmatheus/tech-challenge/src/entities"

type PedidoUseCases interface {
	CriarPedido(entities.Pedido) (entities.Pedido, error)
	RecuperarPedidos() ([]entities.Pedido, error)
	AtualizarStatus(id int, status string) error
}

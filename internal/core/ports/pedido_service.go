package ports

import "github.com/gomesmatheus/tech-challenge/internal/core/domain"

type PedidoService interface {
	CriarPedido(domain.Pedido) (domain.Pedido, error)
	RecuperarPedidos() ([]domain.Pedido, error)
	AtualizarStatus(id int, status string) error
}

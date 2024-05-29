package ports

import "github.com/gomesmatheus/tech-challenge/internal/core/domain"

type PedidoRepository interface {
	CriarPedido(domain.Pedido) (domain.Pedido, error)
	RecuperarPedidos() ([]domain.Pedido, error)
	AtualizarStatus(id int, status string) error
}

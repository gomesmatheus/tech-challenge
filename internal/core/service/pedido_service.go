package service

import (
	"github.com/gomesmatheus/tech-challenge/internal/core/domain"
	"github.com/gomesmatheus/tech-challenge/internal/core/ports"
)

type pedidoService struct {
   repository ports.PedidoRepository 
}

func NewPedidoService(pedidoRepository ports.PedidoRepository) *pedidoService{
    return &pedidoService{
        repository: pedidoRepository,
    }
}

func (service *pedidoService) CriarPedido(p domain.Pedido) (domain.Pedido, error) {
    return service.repository.CriarPedido(p)
}
func (service *pedidoService) RecuperarPedidos() ([]domain.Pedido, error) {
    return service.RecuperarPedidos()
}

func (service *pedidoService) AtualizarStatus(id int, status string) (error) {
    return service.AtualizarStatus(id, status)
}

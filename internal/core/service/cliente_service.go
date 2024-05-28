package service

import (
	"github.com/gomesmatheus/tech-challenge/internal/core/domain"
	"github.com/gomesmatheus/tech-challenge/internal/core/ports"
)

type clienteService struct {
   repository ports.ClienteRepository 
}

func NewClienteService(clienteRepository ports.ClienteRepository) *clienteService{
    return &clienteService{
        repository: clienteRepository,
    }
}

func (service *clienteService) Cadastrar(cliente domain.Cliente) (domain.Cliente, error) {
    // adicionar validação
    return service.repository.Cadastrar(cliente)
}

func (service *clienteService) Recuperar(cpf int64) (domain.Cliente, error) {
    // adicionar validação
    return service.repository.Recuperar(cpf)
}


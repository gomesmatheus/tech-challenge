package service

import (
	"errors"

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
    if !isClienteValido(cliente) {
        return cliente, errors.New("Cliente inválido")
    }

    return service.repository.Cadastrar(cliente)
}

func (service *clienteService) Identificar(cpf int64) (domain.Cliente, error) {
    // adicionar validação de cpf
    return service.repository.Recuperar(cpf)
}

func isClienteValido(cliente domain.Cliente) bool {
    return cliente.Cpf != 0 && cliente.Nome != "" && cliente.Email != ""
}


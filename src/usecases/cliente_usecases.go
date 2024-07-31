package usecases

import (
	"errors"

	"github.com/gomesmatheus/tech-challenge/src/entities"
	"github.com/gomesmatheus/tech-challenge/src/interfaces"
)

type clienteUseCases struct {
   gateway interfaces.ClienteGateway 
}

func NewClienteUseCases(clienteGateway interfaces.ClienteGateway) *clienteUseCases{
    return &clienteUseCases{
        gateway: clienteGateway,
    }
}

func (usecase *clienteUseCases) Cadastrar(cliente entities.Cliente) (entities.Cliente, error) {
    if !isClienteValido(cliente) {
        return cliente, errors.New("Cliente inválido")
    }

    return usecase.gateway.Cadastrar(cliente)
}

func (usecase *clienteUseCases) Identificar(cpf int64) (entities.Cliente, error) {
    // adicionar validação de cpf
    return usecase.gateway.Recuperar(cpf)
}

func isClienteValido(cliente entities.Cliente) bool {
    return cliente.Cpf != 0 && cliente.Nome != "" && cliente.Email != ""
}


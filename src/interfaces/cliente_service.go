package interfaces

import "github.com/gomesmatheus/tech-challenge/src/entities"

type ClienteUseCases interface {
    Cadastrar(entities.Cliente) (entities.Cliente, error)
    Identificar(cpf int64) (entities.Cliente, error)
}

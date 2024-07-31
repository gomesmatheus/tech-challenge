package interfaces

import "github.com/gomesmatheus/tech-challenge/src/entities"

type ClienteGateway interface {
    Cadastrar(entities.Cliente) (entities.Cliente, error)
    Recuperar(cpf int64) (entities.Cliente, error)
}

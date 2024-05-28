package ports

import "github.com/gomesmatheus/tech-challenge/internal/core/domain"

type ClienteService interface {
    Cadastrar(domain.Cliente) (domain.Cliente, error)
    Recuperar(cpf int64) (domain.Cliente, error)
}

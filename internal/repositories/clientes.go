package repositories

import (
	"context"
	"fmt"

	"github.com/gomesmatheus/tech-challenge/internal/core/domain"
)

func (repo *postgresDb) Cadastrar(cliente domain.Cliente) (domain.Cliente, error) {
    _, err := repo.db.Exec(context.Background(), "INSERT INTO clientes (cpf, nome, email) VALUES ($1, $2, $3)", cliente.Cpf, cliente.Nome, cliente.Email)
    if err != nil {
        fmt.Println("Erro ao inserir cliente na base de dados", err)
    }
    return cliente, err
}

func (repo *postgresDb) Recuperar(cpf int64) (domain.Cliente, error) {
    var cliente domain.Cliente
    err := repo.db.QueryRow(context.Background(), "SELECT cpf, nome, email FROM clientes WHERE cpf = $1", cpf).Scan(&cliente.Cpf, &cliente.Nome, &cliente.Email)
    if err != nil {
        fmt.Println("Erro buscando por cpf", cpf, err)
    }
    
    return cliente, err
}


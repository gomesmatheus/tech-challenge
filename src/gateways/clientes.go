package gateways

import (
	"context"
	"fmt"

	"github.com/gomesmatheus/tech-challenge/src/entities"
)

func (repo *DbConnection) Cadastrar(cliente entities.Cliente) (entities.Cliente, error) {
    _, err := repo.Db.Exec(context.Background(), "INSERT INTO clientes (cpf, nome, email) VALUES ($1, $2, $3)", cliente.Cpf, cliente.Nome, cliente.Email)
    if err != nil {
        fmt.Println("Erro ao inserir cliente na base de dados", err)
    }
    return cliente, err
}

func (repo *DbConnection) Recuperar(cpf int64) (entities.Cliente, error) {
    var cliente entities.Cliente
    err := repo.Db.QueryRow(context.Background(), "SELECT cpf, nome, email FROM clientes WHERE cpf = $1", cpf).Scan(&cliente.Cpf, &cliente.Nome, &cliente.Email)
    if err != nil {
        fmt.Println("Erro buscando por cpf", cpf, err)
    }
    
    return cliente, err
}


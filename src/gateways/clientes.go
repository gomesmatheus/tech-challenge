package gateways

import (
	"context"
	"fmt"
    "time"
    "strconv"
    "encoding/json"

	"github.com/gomesmatheus/tech-challenge/src/entities"
)

func (repo *DbConnection) Cadastrar(cliente entities.Cliente) (entities.Cliente, error) {
    _, err := repo.Db.Exec(context.Background(), "INSERT INTO clientes (cpf, nome, email) VALUES ($1, $2, $3)", cliente.Cpf, cliente.Nome, cliente.Email)
    if err != nil {
        fmt.Println("Erro ao inserir cliente na base de dados", err)
    }

    jsonData, _ := json.Marshal(cliente)

    var ctx = context.Background()
    err2 := repo.Redis.Set(ctx, strconv.FormatInt(cliente.Cpf, 10), jsonData, 72*time.Hour).Err()
    if err2 != nil {
        fmt.Printf("Could not set value: %v\n", err2)
    }

    return cliente, err
}

func (repo *DbConnection) Recuperar(cpf int64) (entities.Cliente, error) {
    var cliente entities.Cliente
    var err error
    var ctx = context.Background()
    val, err2 := repo.Redis.Get(ctx, strconv.FormatInt(cpf, 10)).Result()
    if err2 != nil {
        fmt.Printf("Could not get value: %v\n", err2)
        err := repo.Db.QueryRow(context.Background(), "SELECT cpf, nome, email FROM clientes WHERE cpf = $1", cpf).Scan(&cliente.Cpf, &cliente.Nome, &cliente.Email)
        if err != nil {
            fmt.Println("Erro buscando por cpf", cpf, err)
        }
    }

    err = json.Unmarshal([]byte(val), &cliente)
    if err != nil {
        return cliente, err
    }
    
    return cliente, err
}


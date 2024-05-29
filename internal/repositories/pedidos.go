package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/gomesmatheus/tech-challenge/internal/core/domain"
)

// status recebido, em preparação, pronto e finalizado

func (repo *postgresDb) CriarPedido(p domain.Pedido) (domain.Pedido, error) {
    var idPedido int
    err := repo.db.QueryRow(context.Background(), "INSERT INTO pedidos (cliente_cpf, status, data, metodo_pagamento) VALUES ($1, $2, $3, $4) RETURNING id", p.Cpf, "Recebido", time.Now(), p.MetodoPagamento).Scan(&idPedido)
    if err != nil {
        fmt.Println("Erro ao inserir pedido na base de dados", err)
    }

    for _, pp := range p.Produtos {
        _, err := repo.db.Exec(context.Background(), "INSERT INTO produto_pedido (produto_id, pedido_id, quantidade, observacao) values ($1, $2, $3, $4)", pp.ProdutoId, idPedido, pp.Quantidade, pp.Observacao)
        if err != nil {
            fmt.Println("Erro ao inserir pedido na base de dados", err)
            return p, err 
        }
    }

    return p, err
}

func (repo *postgresDb) RecuperarPedidos() ([]domain.Pedido, error) {
    var pedidos []domain.Pedido
    rows, err := repo.db.Query(context.Background(), "SELECT id, cliente_cpf, status, metodo_pagamento FROM pedidos")
    defer rows.Close()
    if err != nil {
        fmt.Println("Erro ao recuperar pedidos")
        fmt.Println(err)
        return nil, err
    }
    
    for rows.Next() {
        var p domain.Pedido
        if err = rows.Scan(&p.Id, &p.Cpf, &p.Status, &p.MetodoPagamento); err != nil {
            fmt.Println("Erro fazendo scanning de pedido")
            fmt.Println(err)
            return nil, err
        }
        pedidos = append(pedidos, p)
    }
    
    return pedidos, err
}


func (repo *postgresDb) AtualizarStatus(idPedido int, status string) (error) {
    _, err := repo.db.Exec(context.Background(), "UPDATE pedidos SET status = $1 WHERE id = $2", status, idPedido)
    if err != nil {
        fmt.Println("Erro ao trocar status do pedido na base de dados", err)
    }
    return err
}



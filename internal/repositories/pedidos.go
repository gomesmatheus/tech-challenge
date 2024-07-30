package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/gomesmatheus/tech-challenge/internal/core/domain"
)

type PedidoRow struct {
	Id                int
	Cpf               int64
	Status            string
	MetodoPagamento   string
	PagamentoAprovado bool
	ProdutoId         int
	Quantidade        int
	Observacao        string
}

const (
	QUERY_PEDIDOS = `
        SELECT
            A.id,
            A.cliente_cpf,
            A.status,
            A.metodo_pagamento,
						A.pagamento_aprovado,
            B.produto_id,
            B.quantidade,
            B.observacao
        FROM pedidos A
        INNER JOIN produto_pedido B ON A.id = B.pedido_id;
    `
)

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
	rows, err := repo.db.Query(context.Background(), QUERY_PEDIDOS)
	if err != nil {
		fmt.Println("Erro ao recuperar pedidos:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r PedidoRow
		if err = rows.Scan(&r.Id, &r.Cpf, &r.Status, &r.MetodoPagamento, &r.PagamentoAprovado, &r.ProdutoId, &r.Quantidade, &r.Observacao); err != nil {
			fmt.Println("Erro fazendo scanning de pedido:", err)
			return nil, err
		}

		pedidoJaExiste := false
		for i, p := range pedidos {
			if p.Id == r.Id {
				pedidoJaExiste = true
				pedidos[i].Produtos = append(p.Produtos, domain.ProdutoPedido{
					ProdutoId:  r.ProdutoId,
					Quantidade: r.Quantidade,
					Observacao: r.Observacao,
				})
			}
		}

		if !pedidoJaExiste {
			pedidos = append(pedidos, domain.Pedido{
				Id:                r.Id,
				Cpf:               r.Cpf,
				Status:            r.Status,
				MetodoPagamento:   r.MetodoPagamento,
				PagamentoAprovado: r.PagamentoAprovado,
				Produtos: []domain.ProdutoPedido{
					{
						ProdutoId:  r.ProdutoId,
						Quantidade: r.Quantidade,
						Observacao: r.Observacao,
					},
				},
			})
		}
	}

	return pedidos, err
}

func (repo *postgresDb) AtualizarStatus(idPedido int, status string) error {
	_, err := repo.db.Exec(context.Background(), "UPDATE pedidos SET status = $1 WHERE id = $2", status, idPedido)
	if err != nil {
		fmt.Println("Erro ao trocar status do pedido na base de dados", err)
	}
	return err
}

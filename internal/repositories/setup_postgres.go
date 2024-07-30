package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

const (
	createTables = `
    CREATE TABLE IF NOT EXISTS clientes (
        cpf BIGINT PRIMARY KEY,
        nome VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE
    );

    CREATE TABLE IF NOT EXISTS categoria_produtos (
        id SERIAL PRIMARY KEY,
        descricao VARCHAR(255) NOT NULL UNIQUE
    );

    INSERT INTO categoria_produtos (descricao) VALUES ('Lanche'), ('Acompanhamento'), ('Bebida'), ('Sobremesa') ON CONFLICT (descricao) DO NOTHING;
    SELECT * FROM categoria_produtos;

    CREATE TABLE IF NOT EXISTS produtos (
        id SERIAL PRIMARY KEY,
        categoria_id INTEGER NOT NULL,
        nome VARCHAR(255) NOT NULL UNIQUE,
        descricao VARCHAR(255) NOT NULL,
        preco FLOAT NOT NULL,
        tempo_de_preparo_minutos INTEGER NOT NULL,

        CONSTRAINT fk_categoria_id FOREIGN KEY(categoria_id) REFERENCES categoria_produtos(id)
    );

    CREATE TABLE IF NOT EXISTS pedidos (
        id SERIAL PRIMARY KEY,
        cliente_cpf BIGINT,
        status VARCHAR(255),
        data TIMESTAMP,
        metodo_pagamento VARCHAR(255),
        pagamento_aprovado BOOLEAN DEFAULT FALSE,

        CONSTRAINT fk_cpf FOREIGN KEY(cliente_cpf) REFERENCES clientes(cpf)
    );

    CREATE TABLE IF NOT EXISTS produto_pedido (
        produto_id INTEGER NOT NULL,
        pedido_id INTEGER NOT NULL,
        quantidade INTEGER NOT NULL,
        observacao VARCHAR,

        PRIMARY KEY (produto_id, pedido_id),
        CONSTRAINT fk_produto FOREIGN KEY (produto_id) REFERENCES produtos(id),
        CONSTRAINT fk_pedido FOREIGN KEY (pedido_id) REFERENCES pedidos(id)
    );
    `
)

type postgresDb struct {
	db *pgx.Conn
}

func NewPostgresDb(url string) (*postgresDb, error) {
	config, err := pgx.ParseConfig(url)
	if err != nil {
		fmt.Println("Error parsing config", err)
		return nil, err
	}

	db, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		fmt.Println("Error creating database connection", err)
		return nil, err
	}
	// setup create table
	if _, err := db.Exec(context.Background(), createTables); err != nil {
		fmt.Println("Error creating table Persons", err)
		return nil, err
	}

	return &postgresDb{
		db: db,
	}, nil
}

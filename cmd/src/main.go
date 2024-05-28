package main

import (
	"log"
	"net/http"

	"github.com/gomesmatheus/tech-challenge/internal/core/service"
	"github.com/gomesmatheus/tech-challenge/internal/handlers"
	"github.com/gomesmatheus/tech-challenge/internal/repositories"
)


func main() {
    postgresDb, err := repositories.NewPostgresDb("postgres://postgres:123@postgres-db:5432/postgres")
    if err != nil {
        log.Fatalf("Error initializing database: %v", err)
    }

    clienteService := service.NewClienteService(postgresDb)
    clienteHandler := handlers.NewClienteHandler(clienteService)

    produtoService := service.NewProdutoService(postgresDb)
    produtoHandler := handlers.NewProdutoHandler(produtoService)


    http.HandleFunc("/cliente", clienteHandler.CriacaoRoute)
    http.HandleFunc("/cliente/", clienteHandler.IdentificacaoRoute)
    http.HandleFunc("/produto", produtoHandler.CriacaoProdutoRoute)
    http.HandleFunc("/produto/", produtoHandler.RecuperarProdutosRoute)
    
    log.Fatal(http.ListenAndServe(":3333", nil))
}

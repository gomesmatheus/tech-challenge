package main

import (
	"log"
	"net/http"

	"github.com/gomesmatheus/tech-challenge/src/external"
	"github.com/gomesmatheus/tech-challenge/src/controllers"
	"github.com/gomesmatheus/tech-challenge/src/usecases"	
)

func main() {
	dbConnection, err := external.NewDbs()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	clienteUseCases := usecases.NewClienteUseCases(dbConnection)
	clienteHandler := controllers.NewClienteHandler(clienteUseCases)

	produtoUseCases := usecases.NewProdutoUseCases(dbConnection)
	produtoHandler := controllers.NewProdutoHandler(produtoUseCases)

	pedidoUseCases := usecases.NewPedidoUseCases(dbConnection)
	pedidoHandler := controllers.NewPedidoHandler(pedidoUseCases)

	pagamentoUseCases := usecases.NewPagamentoUseCases(dbConnection)
	pagamentoHandler := controllers.NewPagamentoHandler(pagamentoUseCases)

	http.HandleFunc("/cliente", clienteHandler.CriacaoRoute)
	http.HandleFunc("/cliente/", clienteHandler.IdentificacaoRoute)
	http.HandleFunc("/produto", produtoHandler.CriacaoProdutoRoute)
	http.HandleFunc("/produto/", produtoHandler.RecuperarProdutosRoute)
	http.HandleFunc("/pedido", pedidoHandler.CriacaoPedidoRoute)
	http.HandleFunc("/pedido/atualizar/", pedidoHandler.AtualizarPedidoRoute)
	http.HandleFunc("/pagamento/webhooks/", pagamentoHandler.AtualizarPagamentoRoute)

	log.Fatal(http.ListenAndServe(":3333", nil))
}

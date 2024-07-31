package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gomesmatheus/tech-challenge/src/interfaces"
)


type PagamentoHandler struct {
    pagamentoUseCases interfaces.PagamentoUseCases
}

type Pagamento struct {
	Id int `json:"id"`
}

func NewPagamentoHandler(pagamentoUseCases interfaces.PagamentoUseCases) *PagamentoHandler {
    return &PagamentoHandler{
        pagamentoUseCases: pagamentoUseCases,
    }
}

func (c *PagamentoHandler) AtualizarPagamentoRoute(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        var pagamento Pagamento
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			fmt.Println("Error parsing request body")
			w.WriteHeader(400)
			w.Write([]byte("400 bad request"))
			return
		}
		json.Unmarshal(body, &pagamento)

        c.pagamentoUseCases.AtualizarPagamento(pagamento.Id, true)
        
        w.WriteHeader(200)
        w.Write([]byte("Pagamento confirmado!"))
    }

    return
}

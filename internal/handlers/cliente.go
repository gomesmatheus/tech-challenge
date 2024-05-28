package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gomesmatheus/tech-challenge/internal/core/domain"
	"github.com/gomesmatheus/tech-challenge/internal/core/ports"
)


type ClienteHandler struct {
    clienteService ports.ClienteService
}

func NewClienteHandler(clienteService ports.ClienteService) *ClienteHandler {
    return &ClienteHandler{
        clienteService: clienteService,
    }
}

func (c *ClienteHandler) CriacaoRoute(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        body, err := io.ReadAll(r.Body)
        defer r.Body.Close()
        if err != nil {
            fmt.Println("Error parsing request body")
            w.WriteHeader(400)
            w.Write([]byte("400 bad request"))
            return
        }

        var cliente domain.Cliente
        json.Unmarshal(body, &cliente)
        fmt.Println(cliente)

        cliente, err = c.clienteService.Cadastrar(cliente)
        if err != nil {
            fmt.Println("Erro ao cadastrar o cliente", err)
            w.WriteHeader(500)
            w.Write([]byte("Erro ao cadastrar o cliente"))
            return
        }
        
        w.WriteHeader(201)
        w.Write([]byte("Cliente inserido"))
    }

    return
}

// func (c *ClienteHandler) IdentificacaoRoute(w http.ResponseWriter, r *http.Request) {
//     if r.Method == "GET" {

//         var cliente domain.Cliente
//         json.Unmarshal(body, &cliente)
//         fmt.Println(cliente)

//         cliente, err = c.clienteService.Cadastrar(cliente)
//         if err != nil {
//             fmt.Println("Erro ao cadastrar o cliente", err)
//             w.WriteHeader(500)
//             w.Write([]byte("Erro ao cadastrar o cliente"))
//             return
//         }
        
//         w.WriteHeader(201)
//         w.Write([]byte("Cliente inserido"))
//     }

//     return
// }


package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gomesmatheus/tech-challenge/internal/core/domain"
	"github.com/gomesmatheus/tech-challenge/internal/core/ports"
)


type PedidoHandler struct {
    pedidoService ports.PedidoService
}

func NewPedidoHandler(pedidoService ports.PedidoService) *PedidoHandler {
    return &PedidoHandler{
        pedidoService: pedidoService,
    }
}

func (c *PedidoHandler) CriacaoPedidoRoute(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        body, err := io.ReadAll(r.Body)
        defer r.Body.Close()
        if err != nil {
            fmt.Println("Error parsing request body")
            w.WriteHeader(400)
            w.Write([]byte("400 bad request"))
            return
        }

        var pedido domain.Pedido
        json.Unmarshal(body, &pedido)
        fmt.Println(pedido)

        pedido, err = c.pedidoService.CriarPedido(pedido)
        if err != nil {
            fmt.Println("Erro ao cadastrar o pedido", err)
            w.WriteHeader(500)
            w.Write([]byte("Erro ao cadastrar o pedido"))
            return
        }
        
        w.WriteHeader(201)
        w.Write([]byte("Pedido inserido"))
    } else if r.Method == "GET" {
        pedidos, err := c.pedidoService.RecuperarPedidos()
        if err != nil {
            fmt.Println("Erro ao recuperar pedidos", err)
            w.WriteHeader(500)
            w.Write([]byte("Erro ao recuperar pedidos"))
            return
        }
        response, _ := json.Marshal(pedidos)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(200)
        w.Write(response)
    }


    return
}


func (c *PedidoHandler) AtualizarPedidoRoute(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.ParseInt(strings.Split(r.URL.Path, "/")[2], 10, 64)
    if r.Method == "PATCH" {
        status := "teste"
        err := c.pedidoService.AtualizarStatus(int(id), status)
        if err != nil {
            fmt.Println("Erro ao atualizar o pedido", err)
            w.WriteHeader(500)
            w.Write([]byte("Erro ao atualizar o pedido"))
            return
        }
        
        w.WriteHeader(201)
        w.Write([]byte("Pedido atualizado"))
    }
}



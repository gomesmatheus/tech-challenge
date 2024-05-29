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


type ProdutoHandler struct {
    produtoService ports.ProdutoService
}

func NewProdutoHandler(produtoService ports.ProdutoService) *ProdutoHandler {
    return &ProdutoHandler{
        produtoService: produtoService,
    }
}

func (c *ProdutoHandler) CriacaoProdutoRoute(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        body, err := io.ReadAll(r.Body)
        defer r.Body.Close()
        if err != nil {
            fmt.Println("Error parsing request body")
            w.WriteHeader(400)
            w.Write([]byte("400 bad request"))
            return
        }

        var produto domain.Produto
        json.Unmarshal(body, &produto)
        fmt.Println(produto)

        produto, err = c.produtoService.CriarProduto(produto)
        if err != nil {
            fmt.Println("Erro ao cadastrar o produto", err)
            w.WriteHeader(500)
            w.Write([]byte("Erro ao cadastrar o produto"))
            return
        }
        
        w.WriteHeader(201)
        w.Write([]byte("Produto inserido"))
    }

    return
}

func (c *ProdutoHandler) RecuperarProdutosRoute(w http.ResponseWriter, r *http.Request) {
    fmt.Println("bateu na rota de recuperar produtos")
    id, err := strconv.ParseInt(strings.Split(r.URL.Path, "/")[2], 10, 64)

    if r.Method == "GET" {
        categoriaId := id
        if err != nil {
            fmt.Println(err)
        }

        produtos, err := c.produtoService.RecuperarProdutos(int(categoriaId))
        if err != nil {
            w.WriteHeader(404)
            w.Write([]byte(fmt.Sprintf("Erro ao recuperar produtos com categoria_id %d", categoriaId)))
            return
        }
        response, _ := json.Marshal(produtos)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(200)
        w.Write(response)
    } else if r.Method == "PUT" {
        body, err := io.ReadAll(r.Body)
        defer r.Body.Close()
        if err != nil {
            fmt.Println("Error parsing request body")
            w.WriteHeader(400)
            w.Write([]byte("400 bad request"))
            return
        }

        var produto domain.Produto
        json.Unmarshal(body, &produto)
        err = c.produtoService.AtualizarProduto(int(id), produto)
        if err != nil {
            fmt.Println("Erro ao atualizar produto")
            w.WriteHeader(500)
            w.Write([]byte("500 Erro ao atualizar produto"))
            return
        }

    } else if r.Method == "DELETE" {
        err := c.produtoService.DeletarProduto(int(id)) 
        if err != nil {
            fmt.Println("Erro ao deletar produto")
            w.WriteHeader(500)
            w.Write([]byte("500 Erro ao deletar produto"))
            return
        }
    }

    return
}


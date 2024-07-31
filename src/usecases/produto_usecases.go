package usecases

import (
	"errors"

	"github.com/gomesmatheus/tech-challenge/src/entities"
	"github.com/gomesmatheus/tech-challenge/src/interfaces"
)

type produtoUseCases struct {
   gateway interfaces.ProdutoGateway 
}

func NewProdutoUseCases(produtoGateway interfaces.ProdutoGateway) *produtoUseCases{
    return &produtoUseCases{
        gateway: produtoGateway,
    }
}

func (usecase *produtoUseCases) CriarProduto(p entities.Produto) (entities.Produto, error) {
    if !isProdutoValido(p) {
        return p, errors.New("Produto inválido")
    }

    return usecase.gateway.CriarProduto(p)
}

func (usecase *produtoUseCases) RecuperarProdutos(categoriaId int) ([]entities.Produto, error) {
    return usecase.gateway.RecuperarProdutos(categoriaId)
}

func (usecase *produtoUseCases) AtualizarProduto(id int, p entities.Produto) (error) {
    if !isProdutoValido(p) {
        return errors.New("Produto inválido")
    }

    return usecase.gateway.AtualizarProduto(id, p)
}

func (usecase *produtoUseCases) DeletarProduto(id int) (error) {
    return usecase.gateway.DeletarProduto(id)
}

func isProdutoValido(p entities.Produto) bool {
    return p.Nome != "" && p.Preco != 0 && p.Descricao != "" && p.CategoriaId != 0 && p.TempoDePreparo != 0 
}


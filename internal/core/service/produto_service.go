package service

import (
	"errors"

	"github.com/gomesmatheus/tech-challenge/internal/core/domain"
	"github.com/gomesmatheus/tech-challenge/internal/core/ports"
)

type produtoService struct {
   repository ports.ProdutoRepository 
}

func NewProdutoService(produtoRepository ports.ProdutoRepository) *produtoService{
    return &produtoService{
        repository: produtoRepository,
    }
}

func (service *produtoService) CriarProduto(p domain.Produto) (domain.Produto, error) {
    if !isProdutoValido(p) {
        return p, errors.New("Produto inválido")
    }

    return service.repository.CriarProduto(p)
}

func (service *produtoService) RecuperarProdutos(categoriaId int) ([]domain.Produto, error) {
    return service.repository.RecuperarProdutos(categoriaId)
}

func (service *produtoService) AtualizarProduto(id int, p domain.Produto) (error) {
    if !isProdutoValido(p) {
        return errors.New("Produto inválido")
    }

    return service.repository.AtualizarProduto(id, p)
}

func (service *produtoService) DeletarProduto(id int) (error) {
    return service.repository.DeletarProduto(id)
}

func isProdutoValido(p domain.Produto) bool {
    return p.Nome != "" && p.Preco != 0 && p.Descricao != "" && p.CategoriaId != 0 && p.TempoDePreparo != 0 
}


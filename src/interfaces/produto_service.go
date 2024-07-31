package interfaces

import "github.com/gomesmatheus/tech-challenge/src/entities"

type ProdutoUseCases interface {
    CriarProduto(p entities.Produto) (entities.Produto, error)
    RecuperarProdutos(categoriaId int) ([]entities.Produto, error)
    AtualizarProduto(id int, p entities.Produto) (error)
    DeletarProduto(id int) (error)
}

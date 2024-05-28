package ports

import "github.com/gomesmatheus/tech-challenge/internal/core/domain"

type ProdutoRepository interface {
    CriarProduto(p domain.Produto) (domain.Produto, error)
    RecuperarProdutos(categoriaId int) ([]domain.Produto, error)
    AtualizarProduto(id int, p domain.Produto) (error)
    DeletarProduto(id int) (error)
}

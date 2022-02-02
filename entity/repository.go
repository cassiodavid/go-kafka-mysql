package entity

type ProdutosRepository interface {
	Insert(product Product) error
}

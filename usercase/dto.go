package usercase

type CreateProductInputDto struct {
	Name       string `json:"name"`
	Quantidade string `json:"quantidade"`
}

type CreateProductOutputDto struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Quantidade string `json:"quantidade"`
}

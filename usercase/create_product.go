package usercase

import (
	"github.com/gokafkamysql/entity"
	"github.com/google/uuid"
)

type CreateProduct struct {
	Repository entity.ProdutosRepository
}

func (c CreateProduct) Execute(input CreateProductInputDto) (CreateProductOutputDto, error) {

	product := entity.Product{}
	product.ID = uuid.NewString()
	product.Name = input.Name
	product.Quantidade = input.Quantidade

	err := c.Repository.Insert(product)
	if err != nil {
		return CreateProductOutputDto{}, err
	}

	output := CreateProductOutputDto{}
	output.ID = product.ID
	output.Name = product.Name
	output.Quantidade = product.Quantidade

	return output, nil

}

func RetornoOutput(input CreateProductInputDto) CreateProductOutputDto {
	return CreateProductOutputDto{
		Name:       input.Name,
		Quantidade: input.Quantidade,
	}
}

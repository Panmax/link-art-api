package service

import (
	"encoding/json"
	"link-art-api/application/command"
	"link-art-api/application/representation"
	"link-art-api/domain/model"
	"link-art-api/domain/repository"
)

func CreateProduct(accountId uint, productCommand *command.CreateProductCommand) error {
	picsJson, _ := json.Marshal(productCommand.DetailPics)

	product := model.NewProduct(accountId, productCommand.Name, productCommand.CategoryId, productCommand.Self,
		productCommand.Price, productCommand.Stock, productCommand.Length, productCommand.Width, productCommand.Year,
		productCommand.Material, productCommand.MainPic, string(picsJson), productCommand.Description)
	return model.CreateOne(product)
}

func ListAccountProduct(accountId uint) ([]*representation.ProductRepresentation, error) {
	var productRepresentations []*representation.ProductRepresentation

	products, err := repository.FindAllProductByUser(accountId)
	if err != nil {
		return nil, err
	}

	for _, p := range products {
		productRepresentations = append(productRepresentations, representation.NewProductRepresentation(p))
	}

	return productRepresentations, nil
}

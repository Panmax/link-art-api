package service

import (
	"encoding/json"
	"link-art-api/application/command"
	"link-art-api/domain/model"
	"link-art-api/domain/repository"
)

func CreateProduct(accountId uint, productCommand *command.CreateProductCommand) error {
	picsJson, _ := json.Marshal(productCommand.DetailPics)

	product := model.NewProduct(accountId, productCommand.Name, productCommand.Type, productCommand.Self,
		productCommand.Price, productCommand.Stock, productCommand.Length, productCommand.Width, productCommand.Year,
		productCommand.Material, productCommand.MainPic, string(picsJson), productCommand.Description)
	return model.CreateOne(product)
}

func ListAccountProduct(accountId uint) ([]model.Product, error) {
	return repository.FindAllProductByUser(accountId)
}

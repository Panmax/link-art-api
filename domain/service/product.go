package service

import (
	"encoding/json"
	"link-art-api/application/command"
	"link-art-api/domain/model"
)

func CreateProduct(accountId uint, productCommand *command.CreateProductCommand) error {
	picsJson, _ := json.Marshal(productCommand.DetailsPics)

	product := model.NewProduct(accountId, productCommand.Name, productCommand.Type, productCommand.Self,
		productCommand.Price, productCommand.Stock, productCommand.Size, productCommand.Year,
		productCommand.Material, productCommand.MainPic, string(picsJson), productCommand.Description)
	return model.CreateOne(product)
}

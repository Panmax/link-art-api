package service

import (
	"errors"
	"link-art-api/application/command"
	"link-art-api/application/representation"
	"link-art-api/domain/model"
	"link-art-api/domain/repository"
)

func CreateProduct(accountId uint, productCommand *command.CreateProductCommand) error {
	picsJson, _ := productCommand.GetDetailPicsJson()

	product := model.NewProduct(accountId, productCommand.Name, productCommand.CategoryId, productCommand.Self,
		productCommand.Price, productCommand.Stock, productCommand.Length, productCommand.Width, productCommand.Year,
		productCommand.Material, productCommand.MainPic, string(picsJson), productCommand.Description)
	return model.CreateOne(product)
}

func UpdateProduct(id uint, accountId *uint, productCommand *command.CreateProductCommand) error {
	picsJson, _ := productCommand.GetDetailPicsJson()
	newProduct := model.NewProduct(*accountId, productCommand.Name, productCommand.CategoryId, productCommand.Self,
		productCommand.Price, productCommand.Stock, productCommand.Length, productCommand.Width, productCommand.Year,
		productCommand.Material, productCommand.MainPic, picsJson, productCommand.Description)

	product, err := repository.FindProduct(id)
	if err != nil {
		return err
	}
	if accountId != nil && product.AccountId != *accountId {
		return errors.New("无权限")
	}
	product.Update(newProduct)

	return model.SaveOne(product)
}

func ListProductByAccount(accountId uint) ([]*representation.ProductRepresentation, error) {
	var productRepresentations []*representation.ProductRepresentation

	products, err := repository.FindAllProductByUser(accountId)
	if err != nil {
		return nil, err
	}

	for _, p := range products {
		productRepresentations = append(productRepresentations, representation.NewProductRepresentation(&p))
	}

	return productRepresentations, nil
}

func GetProduct(id uint) (*representation.ProductRepresentation, error) {
	product, err := repository.FindProduct(id)
	if err != nil {
		return nil, err
	}

	return representation.NewProductRepresentation(product), nil
}

func ShelvesProduct(id uint, accountId *uint) error {
	product, err := repository.FindProduct(id)
	if err != nil {
		return err
	}
	if accountId != nil && product.AccountId != *accountId {
		return errors.New("无权限")
	}
	product.Shelves()
	return model.SaveOne(product)
}

func TakeOffProduct(id uint, accountId *uint) error {
	product, err := repository.FindProduct(id)
	if err != nil {
		return err
	}
	if accountId != nil && product.AccountId != *accountId {
		return errors.New("无权限")
	}
	product.TakeOff()
	return model.SaveOne(product)
}

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

	_, err := repository.FindCategory(productCommand.CategoryId)
	if err != nil {
		return errors.New("作品分类无效")
	}

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
		productRepresentation, err := GetProduct(p.ID)
		if err != nil {
			return nil, err
		}

		productRepresentations = append(productRepresentations, productRepresentation)
	}

	return productRepresentations, nil
}

func GetProduct(id uint) (*representation.ProductRepresentation, error) {
	product, err := repository.FindProduct(id)
	if err != nil {
		return nil, err
	}

	category, err := repository.FindCategory(product.CategoryId)
	if err != nil {
		return nil, err
	}

	return representation.NewProductRepresentation(product, category), nil
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

func ListCategoryByParentId(parentId *uint) ([]*representation.CategoryRepresentation, error) {
	var categoryRepresentations []*representation.CategoryRepresentation

	categories, err := repository.FindAllCategoryByParentId(parentId)
	if err != nil {
		return nil, err
	}
	for _, c := range categories {
		categoryRepresentations = append(categoryRepresentations, representation.NewCategoryRepresentation(&c))
	}

	return categoryRepresentations, nil
}

func ListAllCategory() ([]*representation.CategoryRepresentation, error) {
	categories, err := ListCategoryByParentId(nil)
	if err != nil {
		return nil, err
	}

	err = fillUpChild(categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func fillUpChild(categories []*representation.CategoryRepresentation) error {
	for _, c := range categories {
		children, err := ListCategoryByParentId(&c.ID)
		if err != nil {
			return err
		}
		c.Children = children
		if children != nil {
			return fillUpChild(children)
		}
	}
	return nil
}

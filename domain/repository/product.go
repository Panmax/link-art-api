package repository

import "link-art-api/domain/model"

func FindProduct(id uint) (*model.Product, error) {
	product := &model.Product{}
	err := model.DB.Unscoped().First(product, id).Error
	return product, err
}

func FindAllProductByUser(accountId uint) ([]model.Product, error) {
	var products []model.Product
	err := model.DB.Where("account_id = ?", accountId).Order("created_at desc").Find(&products).Error
	return products, err
}

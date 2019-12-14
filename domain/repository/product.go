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

func FindAllCategoryByParentId(parentId *uint) ([]model.Category, error) {
	var categories []model.Category
	var err error
	if parentId == nil {
		err = model.DB.Where("parent_id is NULL").Order("sort").Find(&categories).Error
	} else {
		err = model.DB.Where("parent_id = ?", parentId).Order("sort").Find(&categories).Error
	}
	return categories, err
}

func FindCategory(id uint) (*model.Category, error) {
	category := &model.Category{}
	err := model.DB.Unscoped().First(category, id).Error
	return category, err
}

func FindAllAuction(accountId uint, auctionType model.AuctionType, status model.AuctionStatus) ([]model.Auction, error) {
	var auctions []model.Auction

	cond := model.DB
	if accountId != 0 {
		cond = cond.Where("account_id = ?", accountId)
	}
	if auctionType != 0 {
		cond = cond.Where("type = ?", auctionType)
	}
	if status != 0 {
		cond = cond.Where("status = ?", status)
	}
	err := cond.Find(&auctions).Error

	return auctions, err
}

func FindAuction(id uint) (*model.Auction, error) {
	auction := &model.Auction{}
	err := model.DB.Unscoped().First(auction, id).Error
	return auction, err
}

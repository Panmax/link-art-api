package repository

import "link-art-api/domain/model"

func FindProduct(id uint) (*model.Product, error) {
	product := &model.Product{}
	err := model.DB.Unscoped().First(product, id).Error
	return product, err
}

func FindAllProduct(args ...interface{}) ([]model.Product, error) {
	var products []model.Product

	cond := model.DB
	if len(args) >= 2 {
		cond = cond.Where(args[0], args[1:]...)
	} else if len(args) >= 1 {
		cond = cond.Where(args[0])
	}

	err := cond.Order("created_at desc").Find(&products).Error

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

func FindAllExhibition(accountId uint, status model.ExhibitionStatus) ([]model.Exhibition, error) {
	var exhibitions []model.Exhibition

	cond := model.DB
	if accountId != 0 {
		cond = cond.Where("account_id = ?", accountId)
	}
	if status != 0 {
		cond = cond.Where("status = ?", status)
	}
	err := cond.Find(&exhibitions).Error

	return exhibitions, err
}

func FindExhibition(id uint) (*model.Exhibition, error) {
	exhibition := &model.Exhibition{}
	err := model.DB.Unscoped().First(exhibition, id).Error
	return exhibition, err
}

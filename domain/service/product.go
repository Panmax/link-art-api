package service

import (
	"container/list"
	"errors"
	"link-art-api/application/command"
	"link-art-api/application/representation"
	"link-art-api/domain/model"
	"link-art-api/domain/repository"
	"time"
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

func ListProduct(accountId uint, keyword string) ([]*representation.ProductRepresentation, error) {

	var products []model.Product
	var err error

	s := "%" + keyword + "%"

	if accountId == 0 && len(keyword) == 0 {
		products, err = repository.FindAllProduct()
	} else if accountId > 0 && len(keyword) == 0 {
		products, err = repository.FindAllProduct("account_id = ? AND status = 1", accountId)
	} else if accountId == 0 && len(keyword) > 0 {
		products, err = repository.FindAllProduct("(name LIKE ? OR description LIKE ?) AND status = 1", s, s)
	} else {
		products, err = repository.FindAllProduct(
			"(name LIKE ? OR description LIKE ?) AND account_id = ? AND status = 1", s, s, accountId)
	}
	if err != nil {
		return nil, err
	}

	productRepresentations := make([]*representation.ProductRepresentation, 0)
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

	artist, err := GetUser(product.AccountId)
	if err != nil {
		return nil, err
	}

	return representation.NewProductRepresentation(product, category, artist), nil
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

func SubmitAuction(accountID uint, submitCommand *command.SubmitAuctionCommand) error {
	startTime := time.Unix(submitCommand.StartTime, 0)

	var items []*model.AuctionItem
	for _, itemCommand := range submitCommand.Items {
		items = append(items, model.NewAuctionItem(itemCommand.ProductID, itemCommand.StartPrice))
	}

	auction := model.NewAuction(accountID, submitCommand.Type, startTime, items)
	return model.CreateOne(auction)
}

func ListAuction(accountId uint, auctionType model.AuctionType, status model.AuctionStatus, action int8) ([]*representation.AuctionRepresentation, error) {
	auctions, err := repository.FindAllAuction(accountId, auctionType, status)

	if err != nil {
		return nil, err
	}

	results := make([]*representation.AuctionRepresentation, 0)
	for _, auction := range auctions {
		auctionRepresentation, err := GetAuction(auction.ID)
		if err != nil {
			return nil, err
		}

		if action != 0 { // 说明不需要过滤
			if auctionRepresentation.Action != action {
				continue
			}
		}
		results = append(results, auctionRepresentation)
	}
	return results, nil
}

func GetAuction(id uint) (*representation.AuctionRepresentation, error) {
	auction, err := repository.FindAuction(id)
	if err != nil {
		return nil, err
	}

	return representation.NewAuctionRepresentation(auction), nil
}

func SubmitExhibition(accountID uint, submitCommand *command.SubmitExhibitionCommand) error {
	startTime := time.Unix(submitCommand.StartTime, 0)
	endTime := time.Unix(submitCommand.EndTime, 0)

	for _, productID := range submitCommand.ProductIDs {
		product, err := repository.FindProduct(productID)
		if err != nil {
			return err
		} else if product.AccountId != accountID {
			return errors.New("加入个展商品有误")
		}
	}

	exhibition := model.NewExhibition(accountID, submitCommand.Title, submitCommand.Description, startTime, endTime, submitCommand.ProductIDs)
	return model.CreateOne(exhibition)
}

func ListExhibition(accountId uint, action int8) ([]*representation.ExhibitionRepresentation, error) {
	exhibitions, err := repository.FindAllExhibition(accountId, model.ExhibitionPassStatus) // 只展示已通过的
	if err != nil {
		return nil, err
	}

	results := make([]*representation.ExhibitionRepresentation, 0)
	for _, exhibition := range exhibitions {
		exhibitionRepresentation, err := GetExhibition(exhibition.ID)
		if err != nil {
			return nil, err
		}

		if action != 0 { // 说明不需要过滤
			if exhibitionRepresentation.Action != action {
				continue
			}
		}
		results = append(results, exhibitionRepresentation)
	}

	return results, nil
}

func GetExhibition(id uint) (*representation.ExhibitionRepresentation, error) {
	exhibition, err := repository.FindExhibition(id)
	if err != nil {
		return nil, err
	}

	artist, err := GetUser(exhibition.AccountId)
	if err != nil {
		return nil, err
	}

	return representation.NewExhibitionRepresentation(exhibition, artist), nil
}

func ListExhibitionProduct(id uint) ([]*representation.ProductRepresentation, error) {
	exhibition, err := repository.FindExhibition(id)
	if err != nil {
		return nil, err
	}

	products := make([]*representation.ProductRepresentation, 0)
	for _, productId := range exhibition.ProductIDs {
		product, err := GetProduct(productId)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func ListCategoryProduct(categoryId uint) ([]*representation.ProductRepresentation, error) {
	results := make([]*representation.ProductRepresentation, 0)

	ids := list.New()
	pushChildCategoryFlatIds(categoryId, ids)
	ids.PushBack(categoryId)

	for p := ids.Front(); p != nil; p = p.Next() {
		products, err := repository.FindAllProduct("category_id = ? AND status = 1", p.Value)
		if err != nil {
			return nil, err
		}

		for _, product := range products {
			productRepresentation, err := GetProduct(product.ID)
			if err != nil {
				return nil, err
			}
			results = append(results, productRepresentation)
		}
	}

	return results, nil
}

func pushChildCategoryFlatIds(categoryId uint, ids *list.List) {
	categories, _ := repository.FindAllCategoryByParentId(&categoryId)
	if categories != nil {
		for _, category := range categories {
			ids.PushBack(category.ID)
			pushChildCategoryFlatIds(category.ID, ids)
		}
	}

	return
}

package service

import (
	"link-art-api/application/representation"
	"link-art-api/domain/model"
	"link-art-api/domain/repository"
)

func ListDiscoveryProduct() ([]*representation.ProductRepresentation, error) {
	products, err := repository.FindAllProduct("status = 1")
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

func ListDiscoveryArtist(accountId uint) ([]*representation.UserRepresentation, error) {
	accounts, err := repository.FindAllAccount("artist = ?", true)
	if err != nil {
		return nil, err
	}

	results := make([]*representation.UserRepresentation, 0)
	for _, account := range accounts {
		if account.ID != accountId && !CheckFollow(accountId, account.ID) {
			user, err := GetUser(account.ID)
			if err != nil {
				return nil, err
			}
			results = append(results, user)
		}
	}

	return results, nil
}

func ListFollowArtistProduct(accountId uint) ([]*representation.ProductRepresentation, error) {
	followerIds, err := listFollowerAccountId(accountId)
	if err != nil {
		return nil, err
	}

	allProducts := make([]model.Product, 0)
	for _, followerId := range followerIds {
		products, err := repository.FindAllProduct("account_id = ? AND status = 1", followerId)
		if err != nil {
			return nil, err
		}
		allProducts = append(allProducts, products...)
	}

	productRepresentations := make([]*representation.ProductRepresentation, 0)
	for _, p := range allProducts {
		productRepresentation, err := GetProduct(p.ID)
		if err != nil {
			return nil, err
		}

		productRepresentations = append(productRepresentations, productRepresentation)
	}

	return productRepresentations, nil
}

package representation

import (
	"encoding/json"
	"link-art-api/domain/model"
)

type ProductRepresentation struct {
	Id          uint                    `json:"id"`
	Name        string                  `json:"name"`
	Category    *CategoryRepresentation `json:"category"`
	Self        bool                    `json:"self"`
	Price       uint                    `json:"price"`
	Stock       int                     `json:"stock"`
	Length      *uint                   `json:"length"`
	Width       *uint                   `json:"width"`
	Year        *string                 `json:"year"`
	Material    string                  `json:"material"`
	MainPic     string                  `json:"main_pic"`
	DetailPics  []string                `json:"detail_pics"`
	Description string                  `json:"description"`
}

func NewProductRepresentation(product *model.Product, category *model.Category) *ProductRepresentation {

	var detailPics []string
	_ = json.Unmarshal([]byte(product.DetailsPicsJson), &detailPics)

	productRepresentation := &ProductRepresentation{
		Id:          product.ID,
		Name:        product.Name,
		Category:    NewCategoryRepresentation(category),
		Self:        product.Self,
		Price:       product.Price,
		Stock:       product.Stock,
		Length:      product.Length,
		Width:       product.Width,
		Year:        product.Year,
		Material:    product.Material,
		MainPic:     product.MainPic,
		DetailPics:  detailPics,
		Description: product.Description,
	}

	return productRepresentation
}

type CategoryRepresentation struct {
	ID       uint                      `json:"id"`
	Name     string                    `json:"name"`
	Children []*CategoryRepresentation `json:"children"`
}

func NewCategoryRepresentation(category *model.Category) *CategoryRepresentation {
	return &CategoryRepresentation{
		ID:       category.ID,
		Name:     category.Name,
		Children: nil,
	}
}

type AuctionRepresentation struct {
	ID        uint   `json:"id"`
	Type      uint8  `json:"type"`
	StartTime uint64 `json:"start_time"`
	Status    uint8  `json:"status"`
}

type AuctionProductRepresentation struct {
	Product    ProductRepresentation `json:"product"`
	StartPrice uint                  `json:"start_price"`
}

type ExhibitionRepresentation struct {
	ID          uint                 `json:"id"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Artist      ArtistRepresentation `json:"artist"`
	StartTime   uint64               `json:"start_time"`
	EndTime     uint64               `json:"end_time"`
	Status      uint8                `json:"status"`
}

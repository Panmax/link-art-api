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

func NewProductRepresentation(product *model.Product) *ProductRepresentation {

	var detailPics []string
	_ = json.Unmarshal([]byte(product.DetailsPicsJson), &detailPics)

	productRepresentation := &ProductRepresentation{
		Id:          product.ID,
		Name:        product.Name,
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
	Id       string
	Name     string
	Children *[]CategoryRepresentation
}

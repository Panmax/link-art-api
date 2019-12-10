package representation

import (
	"encoding/json"
	"link-art-api/domain/model"
)

type ProductRepresentation struct {
	Name        string
	Category    interface{} // TODO
	Self        bool
	Price       uint
	Stock       int
	Length      *uint
	Width       *uint
	Year        *string
	Material    string
	MainPic     string
	DetailPics  []string
	Description string
}

func NewProductRepresentation(product model.Product) *ProductRepresentation {

	var detailPics []string
	_ = json.Unmarshal([]byte(product.DetailsPicsJson), &detailPics)

	productRepresentation := &ProductRepresentation{
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

package representation

import (
	"encoding/json"
	"link-art-api/application/command"
	"link-art-api/domain/model"
	"time"
)

type ProductRepresentation struct {
	Id          uint                    `json:"id"`
	Artist      *ArtistRepresentation   `json:"artist"`
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

	artist := &ArtistRepresentation{
		ID:     product.AccountId,
		Name:   "",  // TODO
		Avatar: nil, // TODO
	}

	productRepresentation := &ProductRepresentation{
		Id:          product.ID,
		Artist:      artist,
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
	ID        uint                `json:"id"`
	Type      model.AuctionType   `json:"type"`
	StartTime int64               `json:"start_time"`
	Status    model.AuctionStatus `json:"status"`
}

func NewAuctionRepresentation(auction *model.Auction) *AuctionRepresentation {
	return &AuctionRepresentation{
		ID:        auction.ID,
		Type:      auction.Type,
		StartTime: auction.StartTime.Unix(),
		Status:    auction.Status,
	}
}

type AuctionProductRepresentation struct {
	Product    ProductRepresentation `json:"product"`
	StartPrice uint                  `json:"start_price"`
}

type ExhibitionRepresentation struct {
	ID          uint                   `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Artist      *ArtistRepresentation  `json:"artist"`
	StartTime   int64                  `json:"start_time"`
	EndTime     int64                  `json:"end_time"`
	Status      model.ExhibitionStatus `json:"status"`
	Action      int8                   `json:"action"`
}

func NewExhibitionRepresentation(exhibition *model.Exhibition) *ExhibitionRepresentation {
	var action int8

	now := time.Now()
	if now.Before(exhibition.StartTime) {
		action = command.ExhibitionActionSoon
	} else if now.After(exhibition.EndTime) {
		action = command.ExhibitionActionEnd
	} else {
		action = command.ExhibitionActionInProcess
	}

	return &ExhibitionRepresentation{
		ID:          exhibition.ID,
		Title:       exhibition.Title,
		Description: exhibition.Description,
		Artist:      nil, // TODO
		StartTime:   exhibition.StartTime.Unix(),
		EndTime:     exhibition.EndTime.Unix(),
		Status:      exhibition.Status,
		Action:      action,
	}
}

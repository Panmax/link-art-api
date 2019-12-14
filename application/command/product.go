package command

import (
	"encoding/json"
	"link-art-api/domain/model"
)

type CreateProductCommand struct {
	Name        string   `json:"name" binding:"required"`
	CategoryId  uint     `json:"category_id" binding:"required"`
	Self        bool     `json:"self" binding:"required"`
	Price       uint     `json:"price" binding:"required"`
	Stock       int      `json:"stock" binding:"required"`
	Length      *uint    `json:"length"`
	Width       *uint    `json:"width"`
	Year        *string  `json:"year"`
	Material    string   `json:"material" binding:"required"`
	MainPic     string   `json:"main_pic" binding:"required"`
	DetailPics  []string `json:"detail_pics" binding:"required"`
	Description string   `json:"description"`
}

func (c *CreateProductCommand) GetDetailPicsJson() (string, error) {
	picsJson, err := json.Marshal(c.DetailPics)
	return string(picsJson), err
}

type SubmitAuctionCommand struct {
	StartTime int64                 `json:"start_time" binding:"required"`
	Type      model.AuctionType     `json:"type" binding:"required"`
	Items     []*AuctionItemCommand `json:"items" binding:"required"`
}

type AuctionItemCommand struct {
	ProductID  uint `json:"product_id" binding:"required"`
	StartPrice uint `json:"start_price" binding:"required"`
}

type SubmitExhibitionCommand struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	StartTime   uint64 `json:"start_time" binding:"required"`
	EndTime     uint64 `json:"end_time" binding:"required"`
	ProductIDs  []uint `json:"product_ids" binding:"required"`
}

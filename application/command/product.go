package command

import "encoding/json"

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

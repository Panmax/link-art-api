package model

type Product struct {
	Model

	AccountId uint `gorm:"not null"`

	Name            string `gorm:"size:64;not null"`
	CategoryId      uint   `gorm:"not null"`
	Self            bool   `gorm:"not null"`
	Price           uint   `gorm:"not null"`
	Stock           int    `gorm:"not null"`
	Length          *uint
	Width           *uint
	Year            *string
	Material        string `gorm:"size:32;not null"`
	MainPic         string `gorm:"size:512;not null"`
	DetailsPicsJson string `gorm:"type:json;not null"`
	Description     string `gorm:"type:text;not null"`
}

func NewProduct(AccountId uint, Name string, CategoryId uint, Self bool, Price uint, Stock int, Length *uint, Width *uint, Year *string,
	Material string, MainPic string, DetailsPicsJson string, Description string) *Product {

	return &Product{
		AccountId:       AccountId,
		Name:            Name,
		CategoryId:      CategoryId,
		Self:            Self,
		Price:           Price,
		Stock:           Stock,
		Length:          Length,
		Width:           Width,
		Year:            Year,
		Material:        Material,
		MainPic:         MainPic,
		DetailsPicsJson: DetailsPicsJson,
		Description:     Description,
	}
}

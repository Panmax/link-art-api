package model

type Product struct {
	Model

	AccountId uint `gorm:"not null"`

	Name            string `gorm:"size:64;not null"`
	Type            uint   `gorm:"not null"`
	Self            bool   `gorm:"not null"`
	Price           uint   `gorm:"not null"`
	Stock           int    `gorm:"not null"`
	Size            string `gorm:"size:32;not null"`
	Year            string `gorm:"size:32;not null"`
	Material        string `gorm:"size:32;not null"`
	MainPic         string `gorm:"size:512;not null"`
	DetailsPicsJson string `gorm:"type:json;not null"`
	Description     string `gorm:"type:text;not null"`
}

func NewProduct(AccountId uint, Name string, Type uint, Self bool, Price uint, Stock int, Size string, Year string,
	Material string, MainPic string, DetailsPicsJson string, Description string) *Product {

	return &Product{
		AccountId:       AccountId,
		Name:            Name,
		Type:            Type,
		Self:            Self,
		Price:           Price,
		Stock:           Stock,
		Size:            Size,
		Year:            Year,
		Material:        Material,
		MainPic:         MainPic,
		DetailsPicsJson: DetailsPicsJson,
		Description:     Description,
	}
}

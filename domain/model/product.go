package model

type Product struct {
	Model

	AccountId uint `gorm:"not null"`

	Name            string `gorm:"size:64;not null"`
	Status          uint   `gorm:"not null"`
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
		Status:          1, // TODO
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

func (p *Product) Update(updatedProduct *Product) {
	p.Name = updatedProduct.Name
	p.CategoryId = updatedProduct.CategoryId
	p.Self = updatedProduct.Self
	p.Price = updatedProduct.Price
	p.Stock = updatedProduct.Stock
	p.Length = updatedProduct.Length
	p.Width = updatedProduct.Width
	p.Year = updatedProduct.Year
	p.Material = updatedProduct.Material
	p.MainPic = updatedProduct.MainPic
	p.DetailsPicsJson = updatedProduct.DetailsPicsJson
	p.Description = updatedProduct.Description
}

func (p *Product) Shelves() {
	p.Status = 1
}

func (p *Product) TakeOff() {
	p.Status = 0
}

type Category struct {
	Model

	Name     string `gorm:"size:64;not null"`
	ParentId *uint
	Sort     int `gorm:"not null"`
}

package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Product struct {
	Model

	AccountId uint `gorm:"not null"`

	Name            string `gorm:"size:64;not null"`
	Status          int8   `gorm:"not null"`
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
		Status:          -1, // TODO 待审核-1，上架1，下架2
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

func (p *Product) Shelves() { // 通过审核 && 上架
	p.Status = 1
}

func (p *Product) TakeOff() { // 下架
	p.Status = 2
}

type Category struct {
	Model

	Name     string `gorm:"size:64;not null"`
	ParentId *uint
	Sort     int `gorm:"not null"`
}

const (
	AuctionLiveType AuctionType = 1
	AuctionTextType AuctionType = 2
)

type AuctionType uint8

const (
	AuctionUnprocessedStatus AuctionStatus = 0
	AuctionPassStatus        AuctionStatus = 1
	AuctionRejectStatus      AuctionStatus = 2
)

type AuctionStatus uint8

type Auction struct {
	Model

	AccountId uint          `gorm:"not null"`
	Type      AuctionType   `gorm:"not null"`
	Status    AuctionStatus `gorm:"not null"`
	StartTime time.Time     `gorm:"not null"`
	Items     AuctionItems  `gorm:"type:json;not null"`
	End       bool          `gorm:"not null"`
}

func NewAuction(accountId uint, auctionType AuctionType, startTime time.Time, items []*AuctionItem) *Auction {
	return &Auction{
		AccountId: accountId,
		Type:      auctionType,
		Status:    AuctionUnprocessedStatus,
		StartTime: startTime,
		Items:     items,
		End:       false,
	}
}

type AuctionItems []*AuctionItem

func (a AuctionItems) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *AuctionItems) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), a)
}

type AuctionItem struct {
	ProductID  uint `json:"product_id"`
	StartPrice uint `json:"start_price"`
}

func NewAuctionItem(productID, startPrice uint) *AuctionItem {
	return &AuctionItem{
		ProductID:  productID,
		StartPrice: startPrice,
	}
}

const (
	ExhibitionUnprocessedStatus ExhibitionStatus = 0
	ExhibitionPassStatus        ExhibitionStatus = 1
	ExhibitionRejectStatus      ExhibitionStatus = 2
)

type ExhibitionStatus uint8

type Exhibition struct {
	Model

	AccountId   uint             `gorm:"not null"`
	Title       string           `gorm:"size:64;not null"`
	Description string           `gorm:"size:512;not null"`
	Status      ExhibitionStatus `gorm:"not null"`
	StartTime   time.Time        `gorm:"not null"`
	EndTime     time.Time        `gorm:"not null"`
	ProductIDs  uintSlice        `gorm:"column:product_ids;type:json;not null"`
}

type uintSlice []uint

func (a uintSlice) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *uintSlice) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), a)
}

func NewExhibition(accountId uint, title string, description string, startTime, endTime time.Time, productIDs []uint) *Exhibition {
	return &Exhibition{
		AccountId:   accountId,
		Title:       title,
		Description: description,
		Status:      ExhibitionUnprocessedStatus,
		StartTime:   startTime,
		EndTime:     endTime,
		ProductIDs:  productIDs,
	}
}

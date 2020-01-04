package model

type Province struct {
	Name       string `gorm:"not null"`
	ProvinceId string `gorm:"not null"`
}

type City struct {
	Name       string `gorm:"not null"`
	CityId     string `gorm:"not null"`
	ProvinceId string `gorm:"not null"`
}

type County struct {
	Name     string `gorm:"not null"`
	CountyId string `gorm:"not null"`
	CityId   string `gorm:"not null"`
}

type Address struct {
	Model

	AccountId  uint   `gorm:"not null"`
	ProvinceId string `gorm:"size:64;not null"`
	CityId     string `gorm:"size:64;not null"`
	CountyId   string `gorm:"size:64;not null"`
	Address    string `gorm:"not null"`
	IsDefault  bool   `gorm:"not null"`
}

func NewAddress(accountId uint, provinceId, cityId, countyId, address string) *Address {
	return &Address{
		AccountId:  accountId,
		ProvinceId: provinceId,
		CityId:     cityId,
		CountyId:   countyId,
		Address:    address,
		IsDefault:  false,
	}
}

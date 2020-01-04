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
	Name       string `gorm:"size:64;not null"`
	Phone      string `gorm:"size:64;not null"`
	ProvinceId string `gorm:"size:64;not null"`
	CityId     string `gorm:"size:64;not null"`
	CountyId   string `gorm:"size:64;not null"`
	Address    string `gorm:"not null"`
}

func NewAddress(accountId uint, name, phone, provinceId, cityId, countyId, address string) *Address {
	return &Address{
		Name:       name,
		Phone:      phone,
		AccountId:  accountId,
		ProvinceId: provinceId,
		CityId:     cityId,
		CountyId:   countyId,
		Address:    address,
	}
}

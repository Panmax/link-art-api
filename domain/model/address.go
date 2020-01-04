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

package representation

import "link-art-api/domain/model"

type ProvinceRepresentation struct {
	Name       string                `json:"name"`
	ProvinceId string                `json:"province_id"`
	Cities     []*CityRepresentation `json:"cities"`
}

type CityRepresentation struct {
	Name       string                  `json:"name"`
	CityId     string                  `json:"city_id"`
	ProvinceId string                  `json:"province_id"`
	Countries  []*CountyRepresentation `json:"counties"`
}

type CountyRepresentation struct {
	Name     string `json:"name"`
	CountyId string `json:"county_id"`
	CityId   string `json:"city_id"`
}

type RegionRepresentation []*ProvinceRepresentation

type AddressRepresentation struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	AccountId    uint   `json:"account"`
	ProvinceId   string `json:"province_id"`
	ProvinceName string `json:"province_name"`
	CityId       string `json:"city_id"`
	CityName     string `json:"city_name"`
	CountyId     string `json:"county_id"`
	CountyName   string `json:"county_name"`
	Address      string `json:"address"`
	IsDefault    bool   `json:"is_default"`
}

func NewAddressRepresentation(address *model.Address, provinceName, cityName, countyName string, IsDefault bool) *AddressRepresentation {
	return &AddressRepresentation{
		Id:           address.ID,
		Name:         address.Name,
		Phone:        address.Phone,
		AccountId:    address.AccountId,
		ProvinceId:   address.ProvinceId,
		ProvinceName: provinceName,
		CityId:       address.CityId,
		CityName:     cityName,
		CountyId:     address.CountyId,
		CountyName:   countyName,
		Address:      address.Address,
		IsDefault:    IsDefault,
	}
}

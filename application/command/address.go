package command

type CreateAddressCommand struct {
	ProvinceId string `json:"province_id" binding:"required"`
	CityId     string `json:"city_id" binding:"required"`
	CountyId   string `json:"county_id" binding:"required"`
	Address    string `json:"address" binding:"required"`
}

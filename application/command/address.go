package command

type CreateAddressCommand struct {
	Name       string `json:"name" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	ProvinceId string `json:"province_id" binding:"required"`
	CityId     string `json:"city_id" binding:"required"`
	CountyId   string `json:"county_id" binding:"required"`
	Address    string `json:"address" binding:"required"`
}

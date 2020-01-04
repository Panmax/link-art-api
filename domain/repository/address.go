package repository

import "link-art-api/domain/model"

func FindAllProvince() ([]model.Province, error) {
	var provinces []model.Province
	err := model.DB.Order("_id").Find(&provinces).Error
	return provinces, err
}

func FindProvince(provinceId string) (*model.Province, error) {
	province := &model.Province{}
	err := model.DB.Where("province_id = ?", provinceId).First(province).Error
	return province, err
}

func FindAllCityByProvinceId(provinceId string) ([]model.City, error) {
	var cities []model.City
	err := model.DB.Where("province_id = ?", provinceId).Order("_id").Find(&cities).Error
	return cities, err
}

func FindCity(cityId string) (*model.City, error) {
	city := &model.City{}
	err := model.DB.Where("city_id = ?", cityId).First(city).Error
	return city, err
}

func FindAllCountyByCityId(cityId string) ([]model.County, error) {
	var counties []model.County
	err := model.DB.Where("city_id = ?", cityId).Order("_id").Find(&counties).Error
	return counties, err
}

func FindCounty(countyId string) (*model.County, error) {
	county := &model.County{}
	err := model.DB.Where("county_id = ?", countyId).First(county).Error
	return county, err
}

func FindAddress(id uint) (*model.Address, error) {
	address := &model.Address{}
	err := model.DB.Unscoped().First(address, id).Error
	return address, err
}

func FindAllAddressByAccount(accountId uint) ([]model.Address, error) {
	var addresses []model.Address
	err := model.DB.Where("account_id = ?", accountId).Order("id desc").Find(&addresses).Error
	return addresses, err
}

func DeleteAddress(id uint) error {
	err := model.DB.Unscoped().Where(
		"id = ?", id).Delete(&model.Address{}).Error
	return err
}

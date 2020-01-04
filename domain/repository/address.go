package repository

import "link-art-api/domain/model"

func FindAllProvince() ([]model.Province, error) {
	var provinces []model.Province
	err := model.DB.Order("_id").Find(&provinces).Error
	return provinces, err
}

func FindAllCityByProvinceId(provinceId string) ([]model.City, error) {
	var cities []model.City
	err := model.DB.Where("province_id = ?", provinceId).Order("_id").Find(&cities).Error
	return cities, err
}

func FindAllCountyByCityId(cityId string) ([]model.County, error) {
	var counties []model.County
	err := model.DB.Where("city_id = ?", cityId).Order("_id").Find(&counties).Error
	return counties, err
}

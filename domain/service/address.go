package service

import (
	"link-art-api/application/representation"
	"link-art-api/domain/repository"
)

func ListRegion() ([]*representation.ProvinceRepresentation, error) {
	regions := make([]*representation.ProvinceRepresentation, 0)
	provinces, err := repository.FindAllProvince()
	if err != nil {
		return nil, err
	}
	for _, province := range provinces {
		regionCites := make([]*representation.CityRepresentation, 0)
		cites, err := repository.FindAllCityByProvinceId(province.ProvinceId)
		if err != nil {
			return nil, err
		}
		for _, city := range cites {
			regionCounties := make([]*representation.CountyRepresentation, 0)
			countries, err := repository.FindAllCountyByCityId(city.CityId)
			if err != nil {
				return nil, err
			}
			for _, county := range countries {
				c := &representation.CountyRepresentation{
					Name:     county.Name,
					CountyId: county.CountyId,
					CityId:   county.CityId,
				}
				regionCounties = append(regionCounties, c)
			}

			c := &representation.CityRepresentation{
				Name:       city.Name,
				CityId:     city.CityId,
				ProvinceId: city.ProvinceId,
				Countries:  regionCounties,
			}
			regionCites = append(regionCites, c)
		}

		p := &representation.ProvinceRepresentation{
			Name:       province.Name,
			ProvinceId: province.ProvinceId,
			Cities:     regionCites,
		}
		regions = append(regions, p)
	}
	return regions, nil
}

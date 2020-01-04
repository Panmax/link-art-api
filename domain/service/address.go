package service

import (
	"encoding/json"
	"link-art-api/application/command"
	"link-art-api/application/representation"
	"link-art-api/domain/model"
	"link-art-api/domain/repository"
	"link-art-api/infrastructure/util/cache"
)

func ListRegion() (representation.RegionPresentation, error) {
	region := representation.RegionPresentation{}

	cacheResult, _ := cache.CACHE.Get(cache.RegionCacheKey).Result()
	if cacheResult != "" {
		err := json.Unmarshal([]byte(cacheResult), &region)
		return region, err
	}

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
		region = append(region, p)
	}

	cacheByte, err := json.Marshal(region)
	cache.CACHE.Set(cache.RegionCacheKey, string(cacheByte), 0)
	return region, err
}

func CreateAddress(accountId uint, addressCommand *command.CreateAddressCommand) error {
	address := model.NewAddress(accountId, addressCommand.ProvinceId, addressCommand.CityId, addressCommand.CountyId, addressCommand.Address)
	return model.SaveOne(address)
}

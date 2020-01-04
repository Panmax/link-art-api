package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"link-art-api/application/command"
	"link-art-api/application/representation"
	"link-art-api/domain/model"
	"link-art-api/domain/repository"
	"link-art-api/infrastructure/util/cache"
)

func ListRegion() (representation.RegionRepresentation, error) {
	region := representation.RegionRepresentation{}

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
	address := model.NewAddress(accountId, addressCommand.Name, addressCommand.Phone,
		addressCommand.ProvinceId, addressCommand.CityId, addressCommand.CountyId, addressCommand.Address)
	return model.SaveOne(address)
}

func ListAddress(accountId uint) ([]*representation.AddressRepresentation, error) {
	addressRepresentations := make([]*representation.AddressRepresentation, 0)
	defaultAddressId, _ := cache.CACHE.Get(fmt.Sprintf(cache.DefaultAddressKey, accountId)).Int()

	addresses, err := repository.FindAllAddressByAccount(accountId)
	if err != nil {
		return nil, err
	}
	for _, address := range addresses {
		a, err := GetAddress(address.ID)
		if err != nil {
			return nil, err
		}
		if uint(defaultAddressId) == a.Id {
			a.IsDefault = true
		}
		addressRepresentations = append(addressRepresentations, a)
	}

	return addressRepresentations, nil
}

func GetAddress(id uint) (*representation.AddressRepresentation, error) {
	address, err := repository.FindAddress(id)
	if err != nil {
		return nil, err
	}
	province, err := repository.FindProvince(address.ProvinceId)
	if err != nil {
		return nil, err
	}
	city, err := repository.FindCity(address.CityId)
	if err != nil {
		return nil, err
	}
	county, err := repository.FindCounty(address.CountyId)
	if err != nil {
		return nil, err
	}

	return representation.NewAddressRepresentation(address, province.Name, city.Name, county.Name, false), nil
}

func UpdateAddress(id uint, addressCommand *command.CreateAddressCommand) error {
	address, err := repository.FindAddress(id)
	if err != nil {
		return err
	}
	address.Update(addressCommand.Name, addressCommand.Phone,
		addressCommand.ProvinceId, addressCommand.CityId, addressCommand.CountyId, addressCommand.Address)

	return model.SaveOne(address)
}

func SetDefaultAddress(accountId, addressId uint) error {
	address, err := repository.FindAddress(addressId)
	if err != nil {
		return err
	}
	if address.AccountId != accountId {
		return errors.New("fuck U")
	}

	key := fmt.Sprintf(cache.DefaultAddressKey, accountId)
	cache.CACHE.Set(key, addressId, 0)

	return nil
}

func DeleteAddress(id uint) error {
	return repository.DeleteAddress(id)
}

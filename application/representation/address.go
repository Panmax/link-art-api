package representation

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

type RegionPresentation []*ProvinceRepresentation

package guarddutyfinding

type City struct {
    CityName string `json:"cityName,omitempty"`
}

func (c *City) SetCityName(cityName string) {
    c.CityName = cityName
}

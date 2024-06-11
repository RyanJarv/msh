package guarddutyfinding

type City_1 struct {
    CityName string `json:"cityName,omitempty"`
}

func (c *City_1) SetCityName(cityName string) {
    c.CityName = cityName
}

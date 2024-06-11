package guarddutyfinding

type City_3 struct {
    CityName string `json:"cityName,omitempty"`
}

func (c *City_3) SetCityName(cityName string) {
    c.CityName = cityName
}

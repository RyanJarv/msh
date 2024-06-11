package guarddutyfinding

type City_2 struct {
    CityName string `json:"cityName,omitempty"`
}

func (c *City_2) SetCityName(cityName string) {
    c.CityName = cityName
}

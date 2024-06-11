package guarddutyfinding

type City_4 struct {
    CityName string `json:"cityName"`
}

func (c *City_4) SetCityName(cityName string) {
    c.CityName = cityName
}

package guarddutyfinding

type Country_4 struct {
    CountryName string `json:"countryName"`
}

func (c *Country_4) SetCountryName(countryName string) {
    c.CountryName = countryName
}

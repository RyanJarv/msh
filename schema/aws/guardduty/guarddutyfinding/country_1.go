package guarddutyfinding

type Country_1 struct {
    CountryName string `json:"countryName,omitempty"`
}

func (c *Country_1) SetCountryName(countryName string) {
    c.CountryName = countryName
}

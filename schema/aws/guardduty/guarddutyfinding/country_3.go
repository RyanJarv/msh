package guarddutyfinding

type Country_3 struct {
    CountryName string `json:"countryName,omitempty"`
}

func (c *Country_3) SetCountryName(countryName string) {
    c.CountryName = countryName
}

package guarddutyfinding

type Country_2 struct {
    CountryName string `json:"countryName,omitempty"`
}

func (c *Country_2) SetCountryName(countryName string) {
    c.CountryName = countryName
}

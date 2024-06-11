package guarddutyfinding

type Country struct {
    CountryName string `json:"countryName,omitempty"`
}

func (c *Country) SetCountryName(countryName string) {
    c.CountryName = countryName
}

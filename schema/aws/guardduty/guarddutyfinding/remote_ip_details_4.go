package guarddutyfinding

type RemoteIpDetails_4 struct {
    City City_4 `json:"city"`
    Country Country_4 `json:"country"`
    GeoLocation GeoLocation_1 `json:"geoLocation"`
    Organization Organization_4 `json:"organization"`
    IpAddressV4 string `json:"ipAddressV4"`
}

func (r *RemoteIpDetails_4) SetCity(city City_4) {
    r.City = city
}

func (r *RemoteIpDetails_4) SetCountry(country Country_4) {
    r.Country = country
}

func (r *RemoteIpDetails_4) SetGeoLocation(geoLocation GeoLocation_1) {
    r.GeoLocation = geoLocation
}

func (r *RemoteIpDetails_4) SetOrganization(organization Organization_4) {
    r.Organization = organization
}

func (r *RemoteIpDetails_4) SetIpAddressV4(ipAddressV4 string) {
    r.IpAddressV4 = ipAddressV4
}

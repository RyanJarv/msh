package guarddutyfinding

type RemoteIpDetails_2 struct {
    City City_2 `json:"city,omitempty"`
    Country Country_2 `json:"country,omitempty"`
    GeoLocation GeoLocation `json:"geoLocation,omitempty"`
    Organization Organization_2 `json:"organization,omitempty"`
    IpAddressV4 string `json:"ipAddressV4,omitempty"`
}

func (r *RemoteIpDetails_2) SetCity(city City_2) {
    r.City = city
}

func (r *RemoteIpDetails_2) SetCountry(country Country_2) {
    r.Country = country
}

func (r *RemoteIpDetails_2) SetGeoLocation(geoLocation GeoLocation) {
    r.GeoLocation = geoLocation
}

func (r *RemoteIpDetails_2) SetOrganization(organization Organization_2) {
    r.Organization = organization
}

func (r *RemoteIpDetails_2) SetIpAddressV4(ipAddressV4 string) {
    r.IpAddressV4 = ipAddressV4
}

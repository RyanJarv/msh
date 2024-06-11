package guarddutyfinding

type RemoteIpDetails_1 struct {
    City City_1 `json:"city,omitempty"`
    Country Country_1 `json:"country,omitempty"`
    GeoLocation GeoLocation `json:"geoLocation,omitempty"`
    Organization Organization_1 `json:"organization,omitempty"`
    IpAddressV4 string `json:"ipAddressV4,omitempty"`
}

func (r *RemoteIpDetails_1) SetCity(city City_1) {
    r.City = city
}

func (r *RemoteIpDetails_1) SetCountry(country Country_1) {
    r.Country = country
}

func (r *RemoteIpDetails_1) SetGeoLocation(geoLocation GeoLocation) {
    r.GeoLocation = geoLocation
}

func (r *RemoteIpDetails_1) SetOrganization(organization Organization_1) {
    r.Organization = organization
}

func (r *RemoteIpDetails_1) SetIpAddressV4(ipAddressV4 string) {
    r.IpAddressV4 = ipAddressV4
}

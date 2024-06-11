package guarddutyfinding

type RemoteIpDetails_3 struct {
    City City_3 `json:"city,omitempty"`
    Country Country_3 `json:"country,omitempty"`
    GeoLocation GeoLocation `json:"geoLocation,omitempty"`
    Organization Organization_3 `json:"organization,omitempty"`
    IpAddressV4 string `json:"ipAddressV4,omitempty"`
}

func (r *RemoteIpDetails_3) SetCity(city City_3) {
    r.City = city
}

func (r *RemoteIpDetails_3) SetCountry(country Country_3) {
    r.Country = country
}

func (r *RemoteIpDetails_3) SetGeoLocation(geoLocation GeoLocation) {
    r.GeoLocation = geoLocation
}

func (r *RemoteIpDetails_3) SetOrganization(organization Organization_3) {
    r.Organization = organization
}

func (r *RemoteIpDetails_3) SetIpAddressV4(ipAddressV4 string) {
    r.IpAddressV4 = ipAddressV4
}

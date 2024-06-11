package guarddutyfinding

type RemoteIpDetails struct {
    City City `json:"city,omitempty"`
    Country Country `json:"country,omitempty"`
    GeoLocation GeoLocation `json:"geoLocation,omitempty"`
    Organization Organization `json:"organization,omitempty"`
    IpAddressV4 string `json:"ipAddressV4,omitempty"`
}

func (r *RemoteIpDetails) SetCity(city City) {
    r.City = city
}

func (r *RemoteIpDetails) SetCountry(country Country) {
    r.Country = country
}

func (r *RemoteIpDetails) SetGeoLocation(geoLocation GeoLocation) {
    r.GeoLocation = geoLocation
}

func (r *RemoteIpDetails) SetOrganization(organization Organization) {
    r.Organization = organization
}

func (r *RemoteIpDetails) SetIpAddressV4(ipAddressV4 string) {
    r.IpAddressV4 = ipAddressV4
}

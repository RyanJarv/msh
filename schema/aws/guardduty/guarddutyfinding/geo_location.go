package guarddutyfinding

type GeoLocation struct {
    Lat float64 `json:"lat,omitempty"`
    Lon float64 `json:"lon,omitempty"`
}

func (g *GeoLocation) SetLat(lat float64) {
    g.Lat = lat
}

func (g *GeoLocation) SetLon(lon float64) {
    g.Lon = lon
}

package guarddutyfinding

type GeoLocation_1 struct {
    Lat float64 `json:"lat"`
    Lon float64 `json:"lon"`
}

func (g *GeoLocation_1) SetLat(lat float64) {
    g.Lat = lat
}

func (g *GeoLocation_1) SetLon(lon float64) {
    g.Lon = lon
}

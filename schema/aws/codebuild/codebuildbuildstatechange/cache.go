package codebuildbuildstatechange

type Cache struct {
    Location string `json:"location,omitempty"`
    CacheType string `json:"type"`
}

func (c *Cache) SetLocation(location string) {
    c.Location = location
}

func (c *Cache) SetCacheType(cacheType string) {
    c.CacheType = cacheType
}

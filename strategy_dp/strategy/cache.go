package strategy

type Cache struct {
	Storage      map[string]string
	evictionAlgo EvictionAlgo
	Capacity     int
	MaxCapacity  int
}

func InitCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		Storage:      storage,
		evictionAlgo: e,
		Capacity:     0,
		MaxCapacity:  2,
	}
}

func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) Add(key, value string) {
	if c.Capacity == c.MaxCapacity {
		c.evictionAlgo.evict(c)
		c.Capacity--
	}
	c.Capacity++
	c.Storage[key] = value
}

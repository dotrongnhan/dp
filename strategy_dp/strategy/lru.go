package strategy

import "fmt"

type LRU struct {
}

func (f *LRU) evict(c *Cache) {
	fmt.Println("Evicting by LRU strategy")
}

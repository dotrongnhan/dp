package strategy

import "fmt"

type LFU struct {
}

func (f *LFU) evict(c *Cache) {
	fmt.Println("Evicting by LFU strategy")
}

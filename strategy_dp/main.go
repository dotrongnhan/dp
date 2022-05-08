package main

import "strategy/strategy"

func main() {
	lfu := &strategy.LFU{}
	cache := strategy.InitCache(lfu)
	cache.Add("a", "1")
	cache.Add("b", "3")
	cache.Add("c", "5")
	lru := &strategy.LRU{}
	cache.SetEvictionAlgo(lru)
	cache.Add("d", "2")
	fifo := &strategy.FIFO{}
	cache.SetEvictionAlgo(fifo)
	cache.Add("v", "211")
}

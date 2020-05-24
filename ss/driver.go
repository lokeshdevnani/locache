package main

import "locache"

func get(cache *locache.Locache, str string) {
	val, found := cache.Get(str)
	if found {
		println(val.(string))
	} else {
		println("not found")
	}
}

func main() {
	config := &locache.Config{SegmentCount: 1, PolicyName: "lru", LRUCapacity: 3}
	cache := locache.New(config)

	// get(cache, "hello")
	cache.Put("A", "1")
	cache.Put("B", "2")
	cache.Put("C", "3")
	cache.Put("D", "4")
	cache.Put("E", "5")

	get(cache, "E")
	get(cache, "A")
	get(cache, "B")
	get(cache, "D")

}

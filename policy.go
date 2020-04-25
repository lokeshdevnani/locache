package locache

func Run() {
	capacity := 3
	segment := newCacheSegment()
	var policy Policy = NewLruPolicy(segment, capacity)
	// v := []byte("v")
	entryA := newCacheEntry("K", "V")

	policy.Put(entryA)

	policy.Put(newCacheEntry("A", "1"))
	policy.Put(newCacheEntry("B", "2"))
	z, _ := policy.Get("K")
	println(z.value.(string))
	policy.Put(newCacheEntry("C", "3"))
	// key := "A"
	z, _ = policy.Get("K")
	println(z.value.(string))
	println(z)
}

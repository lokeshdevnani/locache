package locache

type Policy interface {
	Setup()
	Get(CacheKey) (*CacheEntry, bool)
	Put(*CacheEntry) bool
	Remove(key CacheKey) bool
}

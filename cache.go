package locache

type Config struct {
	segmentCount int
}

type Locache struct {
	segments []*cacheSegment
	config   Config
	policy   *Policy
}

type LocacheConfig struct {
	segmentCount int
	policyName   string
}

func NewLocache(config LocacheConfig) *Locache {
	return &Locache{}
}

// CacheKey for cache
type CacheKey interface{}

// CacheValue for cache
type CacheValue interface{}

// CacheEntry encapsulates the key, value and metadata for a cache entry
type CacheEntry struct {
	key   CacheKey
	value CacheValue
}

func newCacheEntry(key CacheKey, value CacheValue) *CacheEntry {
	return &CacheEntry{key, value}
}

func gen() {
	config := &LocacheConfig{segmentCount: 4, policyName: "lru"}
	var cache = NewLocache(config)
}

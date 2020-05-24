package locache

import (
	"errors"

	"github.com/OneOfOne/xxhash"
)

type Config struct {
	SegmentCount int
	PolicyName   string
	LRUCapacity  int
}

type Locache struct {
	segments []Policy
	config   Config
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

func (config *Config) validate() {
	if config.PolicyName == "lru" {
		if config.LRUCapacity <= 0 {
			config.LRUCapacity = 100
		}
	}
}

// New creates a new locache instance with specified config and returns
func New(config *Config) *Locache {
	config.validate()
	segments := make([]Policy, config.SegmentCount)
	for i := range segments {
		segments[i], _ = buildPolicy(config)
	}
	return &Locache{
		segments: segments,
	}
}

// Get returns the value corresponding to the given key if value is found,
// along with a boolean indicating whether it was found.
func (cache *Locache) Get(key CacheKey) (CacheValue, bool) {
	segment := cache.selectSegment(key)
	entry, found := segment.Get(key)
	if found {
		return entry.value, true
	}
	return nil, false
}

// Put stores the given key-value pair to the cache
func (cache *Locache) Put(key CacheKey, value CacheValue) bool {
	segment := cache.selectSegment(key)
	entry := newCacheEntry(key, value)
	return segment.Put(entry)
}

func segmentHash(key CacheKey) uint64 {
	switch k := key.(type) {
	case string:
		bytes := []byte(k)
		return xxhash.Checksum64(bytes)
	case int32:
		return uint64(k)
	case interface{}:
		// return key
		// xxhash.Checksum64([]byte(k))
		//

	}
	return 0
}

func (cache *Locache) selectSegment(key CacheKey) Policy {
	segmentNumber := segmentHash(key)
	return cache.segments[segmentNumber]
}

func buildPolicy(config *Config) (Policy, error) {
	if config.PolicyName == "lru" {
		policy := &LruPolicy{capacity: config.LRUCapacity}
		policy.Setup()
		return policy, nil
	}

	return nil, errors.New("Unsupported Policy: " + config.PolicyName)
}

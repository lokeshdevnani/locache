package locache

import (
	"container/list"
	"sync"
)

type node = list.Element

type cacheSegment struct {
	lock sync.RWMutex
	data map[CacheKey]*list.Element
}

func newCacheSegment() *cacheSegment {
	return &cacheSegment{
		data: make(map[CacheKey]*list.Element),
	}
}

func extractEntry(element *list.Element) *CacheEntry {
	return element.Value.(*CacheEntry)
}

func setEntry(element *list.Element, entry *CacheEntry) {
	element.Value = entry
}

// TODO: add ruby go boundary code as well

type Policy interface {
	Get(CacheKey) (*CacheEntry, bool)
	Put(*CacheEntry) bool
	Remove(key *CacheKey) bool
}

var _ Policy = (*LruPolicy)(nil)

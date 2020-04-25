package locache

import "container/list"

type LruPolicy struct {
	segment  *cacheSegment
	capacity int
	dll      *list.List
}

func NewLruPolicy(cacheSegment *cacheSegment, capacity int) *LruPolicy {
	return &LruPolicy{
		segment:  cacheSegment,
		capacity: capacity,
		dll:      list.New(),
	}
}

func (lru *LruPolicy) Get(key CacheKey) (*CacheEntry, bool) {
	if element, present := lru.segment.data[key]; present {
		lru.dll.MoveToFront(element)
		return extractEntry(element), true
	}
	return nil, false
}

func (lru *LruPolicy) Put(entry *CacheEntry) bool {
	if element, present := lru.segment.data[entry.key]; present {
		setEntry(element, entry)
		lru.dll.MoveToFront(element)
		return true
	}
	element := lru.dll.PushFront(entry)
	lru.segment.data[entry.key] = element

	lru.removeOldIfRequired()
	return false
}

func (lru *LruPolicy) Remove(key *CacheKey) bool {
	if element, present := lru.segment.data[key]; present {
		lru.removeElement(element)
		return true
	}
	return false
}

func (lru *LruPolicy) removeOldIfRequired() {
	for lru.dll.Len() > lru.capacity {
		lru.removeElement(lru.dll.Back())
	}
}

func (lru *LruPolicy) removeElement(element *list.Element) {
	entry := extractEntry(element)
	delete(lru.segment.data, entry.key)
	lru.dll.Remove(element)
}

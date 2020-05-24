package locache

import (
	"container/list"
	"sync"
)

type LruPolicy struct {
	lock     sync.RWMutex
	capacity int
	dll      *list.List
	data     map[CacheKey]*list.Element
}

func NewLruPolicy(capacity int) *LruPolicy {
	return &LruPolicy{
		capacity: capacity,
		dll:      list.New(),
		data:     make(map[CacheKey]*list.Element),
	}
}

func (lru *LruPolicy) Setup() {
	lru.dll = list.New()
	lru.data = make(map[CacheKey]*list.Element)
}

func (lru *LruPolicy) Get(key CacheKey) (*CacheEntry, bool) {
	if element, present := lru.data[key]; present {
		lru.dll.MoveToFront(element)
		return extractEntry(element), true
	}
	return nil, false
}

func (lru *LruPolicy) Put(entry *CacheEntry) bool {
	if element, present := lru.data[entry.key]; present {
		setEntry(element, entry)
		lru.dll.MoveToFront(element)
		return true
	}
	element := lru.dll.PushFront(entry)
	lru.data[entry.key] = element

	lru.removeOldIfRequired()
	return false
}

func (lru *LruPolicy) Remove(key CacheKey) bool {
	if element, present := lru.data[key]; present {
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
	delete(lru.data, entry.key)
	lru.dll.Remove(element)
}

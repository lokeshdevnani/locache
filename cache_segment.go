package locache

import (
	"container/list"
)

type node = list.Element

func extractEntry(element *list.Element) *CacheEntry {
	return element.Value.(*CacheEntry)
}

func setEntry(element *list.Element, entry *CacheEntry) {
	element.Value = entry
}

// TODO: add ruby go boundary code as well

var _ Policy = (*LruPolicy)(nil)

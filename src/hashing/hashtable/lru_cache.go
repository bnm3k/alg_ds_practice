package hashtable

import (
	"container/list"
)

type payload struct {
	value    int
	lruEntry *list.Element
}

//LRUCache is a cache whereby on fully capacity, the least
//recently used value is evicted
type LRUCache struct {
	lruList  *list.List
	kvStore  map[int]payload
	capacity int
}

//NewLRUCache creates and returns a new lruCache with given capacity
//If capacity less than 1, returns capacity 1
func NewLRUCache(capacity int) LRUCache {
	if capacity < 1 {
		capacity = 1
	}
	return LRUCache{
		lruList:  list.New(),
		kvStore:  make(map[int]payload),
		capacity: capacity,
	}
}

//Get returns value of given key or -1 if key not present
// or has already been evicted
func (cache *LRUCache) Get(key int) int {
	entry, isPresent := cache.kvStore[key]
	if isPresent == false {
		return -1
	}
	//push value to head of list then return
	cache.lruList.MoveToFront(entry.lruEntry)

	return entry.value
}

//Put adds both Key and value to cache
//A put counts as an access hence on insertion
//the value is the most-recently-used
func (cache *LRUCache) Put(key int, value int) {
	//add new entry
	entry, isPresent := cache.kvStore[key]
	if isPresent { //is update
		cache.lruList.MoveToFront(entry.lruEntry)
	} else { //new entry
		entry.lruEntry = cache.lruList.PushFront(key)
	}
	entry.value = value
	cache.kvStore[key] = entry

	//if over capacity, get lru key which should be at
	//the end of the lruList and remove from both
	//the lruList and kvStore
	if len(cache.kvStore) > cache.capacity {
		lruKey := cache.lruList.Back()
		key, _ := lruKey.Value.(int)
		delete(cache.kvStore, key)
		cache.lruList.Remove(lruKey)
	}
}

package hashtable

type payloadLFU struct {
	frequency int
	value     int
}

var emptyVal struct{}

type set map[int]struct{}

//LFUCache ...
type LFUCache struct {
	lfuList  []set
	kvStore  map[int]payloadLFU
	capacity int
}

//Constructor ...
func Constructor(capacity int) LFUCache {
	if capacity < 1 {
		capacity = 1
	}
	lfuList := make([]set, 1)
	lfuList[0] = make(map[int]struct{})
	return LFUCache{
		lfuList:  lfuList,
		kvStore:  make(map[int]payloadLFU),
		capacity: capacity,
	}
}

func (cache *LFUCache) updateFrequency(key int, entry payloadLFU) {
	delete(cache.lfuList[entry.frequency], key)
	entry.frequency++
	cache.kvStore[key] = entry
	if entry.frequency == len(cache.lfuList) {
		cache.lfuList = append(cache.lfuList, make(map[int]struct{}))
	}
	cache.lfuList[entry.frequency][key] = emptyVal
}

//Get ...
func (cache *LFUCache) Get(key int) int {
	entry, isPresent := cache.kvStore[key]
	if isPresent == false {
		return -1
	}
	cache.updateFrequency(key, entry)
	return entry.value
}

func (cache *LFUCache) evictExtra() {
	if len(cache.kvStore) >= cache.capacity {
		for _, bucket := range cache.lfuList {
			if len(bucket) == 0 {
				continue
			}
			for keyToEvict := range bucket {
				delete(bucket, keyToEvict)
				delete(cache.kvStore, keyToEvict)
				return
			}
		}
	}
}

//Put ...
func (cache *LFUCache) Put(key int, value int) {
	//add new entry
	entry, isPresent := cache.kvStore[key]
	if isPresent { //is update
		cache.updateFrequency(key, entry)
	} else { //new entry
		cache.evictExtra()
		entry.frequency = 0
		cache.lfuList[0][key] = emptyVal
	}
	entry.value = value
	cache.kvStore[key] = entry
}

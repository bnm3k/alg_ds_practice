package hashtable

import (
	"container/list"
	"errors"
)

/*
WARNING: fails 3rd Test
*/

/*
A single bucket stores all the keys with an equal number of
frequency access. A dictionary [frequencySet] is used for fast
lookup of random keys. Such lookups occur when a key is being used
hence needs to be bumped up to the next bucket. An accompanying linked-list
is used for popping off the least-recently used key. When a key is added to the bucket,
its also added to the head of the linked list. Hence, at any given instance, the tail
of the list counts as the least-recently-used for that bucket. The dictionary's value
is a pointer to the accompanying linked-list node of the key so that during both
random removal and popping the lru-key, the operations are O(1).
Note that adding is also O(1) since it's adding to the head of the list and the dict
*/
type bucket struct {
	frequencySet map[int]*list.Element
	lruList      *list.List
}

func newBucket() *bucket {
	return &bucket{
		frequencySet: make(map[int]*list.Element),
		lruList:      list.New(),
	}
}

func (b *bucket) add(key int) {
	elem := b.lruList.PushFront(key)
	b.frequencySet[key] = elem
}

func (b *bucket) remove(key int) {
	elem, isPresent := b.frequencySet[key]
	if isPresent {
		b.lruList.Remove(elem)
		delete(b.frequencySet, key)
	}

}

func (b *bucket) isEmpty() bool {
	return len(b.frequencySet) == 0
}

func (b *bucket) popLRU() (int, bool) {
	if len(b.frequencySet) == 0 {
		return -1, false
	}
	lastElem := b.lruList.Back()
	key := lastElem.Value.(int)
	b.lruList.Remove(lastElem)
	delete(b.frequencySet, key)
	return key, true
}

/*
LFUCache encompasses both the key-value map and a lfuList that's
used to track the frequencies of use of the keys.
An insertion does not count as use, hence frequency of use at that
point is zero.
On the other hand, an update and get count as a use.
For each use, the key is bumped up to the next bucket. Read section
on bucket to see why this operation is O(1)
Note, bumping up entails deleting from the current bucket and adding
to the next bucket.
On max capacity, the least-frequently-used key plus its value are evicted
However, if multiple keys are in the same frequency bucket, then the
least-recently-used key is evicted. See section on bucket to see how
it keeps track of the lru.
Note however that the eviction operation could potentially be O(f)
where f is the max frequency. In future, in order to mitigate this,
an additional field will be added to the lfuCache struct to keep track
of the index of the bucket where the lfu key should be
*/

type payload struct {
	frequency int
	value     int
}

//LFUCache ...
type LFUCache struct {
	lfuList  []*bucket
	kvStore  map[int]payload
	capacity int
}

//Constructor ...
func Constructor(capacity int) LFUCache {
	if capacity < 0 {
		panic("Invalid capacity")
	}
	lfuList := make([]*bucket, 1)
	lfuList[0] = newBucket()
	return LFUCache{
		lfuList:  lfuList,
		kvStore:  make(map[int]payload),
		capacity: capacity,
	}
}

func (cache *LFUCache) updateFrequency(key int, entry payload) {
	cache.lfuList[entry.frequency].remove(key)
	entry.frequency++
	cache.kvStore[key] = entry
	if entry.frequency == len(cache.lfuList) {
		cache.lfuList = append(cache.lfuList, newBucket())
	}
	cache.lfuList[entry.frequency].add(key)
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

var errorEvicting = errors.New("Error on eviction, incorrect lfuCache state")

func (cache *LFUCache) evictExtra() error {
	if len(cache.kvStore) >= cache.capacity {
		for _, bucket := range cache.lfuList {
			keyToEvict, isNotEmpty := bucket.popLRU()
			if isNotEmpty {
				delete(cache.kvStore, keyToEvict)
				return nil
			}
		}
		return errorEvicting
	}
	return nil
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
		cache.lfuList[0].add(key)
	}
	entry.value = value
	if cache.capacity > 0 {
		cache.kvStore[key] = entry
	}
}

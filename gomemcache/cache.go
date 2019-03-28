package main

import (
	"errors"
	"reflect"
	"sync"
	"unsafe"
)

// Cache represents cache memory.
type Cache struct {
	sync.Mutex // inherit
	curBytes   int
	maxBytes   int
	hashmap    map[string]*Item
	head       Item // indicates first item of LRU list, not pointer type.
	rear       Item // indicates last item of LRU list, not pointer type.
}

// Item represents a value stored in the cache.
type Item struct {
	key   string
	value []byte
	// folowing members using for LRU list.
	next *Item
	prev *Item
}

// Size returns the real size of an item obj hold by system.
func (item *Item) Size() int {
	return int(int(unsafe.Sizeof(*item)) + len(item.key) + len(item.value))
}

// GetCache returns an cache instance with specified size.
func GetCache(maxBytes int) *Cache {
	cache := &Cache{
		curBytes: 0,
		maxBytes: maxBytes,
		hashmap:  make(map[string]*Item),
	}

	cache.head.next = &cache.rear
	cache.rear.prev = &cache.head

	return cache
}

// Set specified <key, value> pair to memory.
func (cache *Cache) Set(key string, value []byte) error {
	// TODO: validate key and value.

	cache.Lock()
	defer cache.Unlock()

	copyValue := make([]byte, len(value))
	copy(copyValue, value)

	i, ok := cache.hashmap[key]
	if ok {
		isEqualValue := reflect.DeepEqual(i.value, copyValue)
		if isEqualValue {
			cache.fresh(key)
			return nil
		}
		cache.delete(key)
	}

	// It's a fresh <key, value> pair obj, put into hashmap and adjust LRU list.
	newItem := &Item{
		key:   key,
		value: copyValue,
	}

	newItemSize := newItem.Size()
	if newItemSize > cache.maxBytes {
		return errors.New("No avalible memory space")
	}

	rearPrev := cache.rear.prev
	rearPrev.next = newItem
	newItem.prev = rearPrev
	cache.rear.prev = newItem
	newItem.next = &cache.rear

	cache.hashmap[key] = newItem
	cache.curBytes += newItemSize

	cache.checkOverflow()

	return nil
}

// Get value by specified key.
func (cache *Cache) Get(key string) (value []byte, ok bool) {
	cache.Lock()
	defer cache.Unlock()

	// TODO: validate key.
	i, ok := cache.hashmap[key]
	if ok == false {
		return nil, false
	}

	// move the item got by you to rear of the LRU list.
	cache.fresh(key)

	// do not return `i.value` directly.
	result := make([]byte, len(i.value))
	copy(result, i.value)

	return result, true
}

// Delete an item in hashmap, also in LRU list.
func (cache *Cache) Delete(key string) {
	cache.Lock()
	defer cache.Unlock()
	cache.delete(key)
}

func (cache *Cache) fresh(key string) {
	i, ok := cache.hashmap[key]
	if ok == false {
		return
	}

	prevItem := i.prev
	nextItem := i.next
	prevItem.next = nextItem
	nextItem.prev = prevItem

	rearPrev := cache.rear.prev
	rearPrev.next = i
	i.prev = rearPrev
	cache.rear.prev = i
	i.next = &cache.rear
}

func (cache *Cache) checkOverflow() {
	for {
		if cache.curBytes < cache.maxBytes {
			break
		}
		cache.pop()
	}
}

func (cache *Cache) pop() {
	firstItem := cache.head.next
	if firstItem == &cache.rear {
		return
	}

	cache.delete(firstItem.key)
}

func (cache *Cache) delete(key string) {
	i, ok := cache.hashmap[key]
	if ok == false {
		return
	}

	memBytes := i.Size()
	cache.curBytes -= memBytes

	// adjust LRU list
	prevItem := i.prev
	nextItem := i.next
	prevItem.next = nextItem
	nextItem.prev = prevItem

	delete(cache.hashmap, key)
}

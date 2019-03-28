package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheSet(t *testing.T) {
	cache := GetCache(180)
	assert.Equal(t, cache.maxBytes, 180)
	assert.Equal(t, cache.curBytes, 0)
	assert.Equal(t, cache.head.next, &cache.rear)
	assert.Equal(t, cache.rear.prev, &cache.head)
	assert.Equal(t, len(cache.hashmap), 0)

	jerryItem := &Item{
		key:   "name",
		value: []byte("jerry"),
	}
	jerryItemSize := jerryItem.Size()

	cache.Set("name", []byte("jerry"))
	assert.Equal(t, cache.curBytes, jerryItemSize)
	assert.Equal(t, cache.curBytes, 65)
	assert.Equal(t, cache.head.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev, &cache.head)
	item1, ok1 := cache.hashmap["name"]
	assert.True(t, ok1)
	assert.Equal(t, string(item1.value), "jerry")
	assert.Equal(t, len(cache.hashmap), 1)

	cache.Set("name", []byte("jerry"))
	assert.Equal(t, cache.curBytes, jerryItemSize)
	assert.Equal(t, cache.curBytes, 65)
	assert.Equal(t, cache.head.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev, &cache.head)
	item2, ok2 := cache.hashmap["name"]
	assert.True(t, ok2)
	assert.Equal(t, string(item2.value), "jerry")
	assert.Equal(t, len(cache.hashmap), 1)

	cache.Set("name", []byte("jerr"))
	assert.Equal(t, cache.curBytes, jerryItemSize-1)
	assert.Equal(t, cache.curBytes, 64)
	assert.Equal(t, cache.head.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev, &cache.head)
	item3, ok3 := cache.hashmap["name"]
	assert.True(t, ok3)
	assert.Equal(t, string(item3.value), "jerr")
	assert.Equal(t, len(cache.hashmap), 1)

	_, ok4 := cache.hashmap["age"]
	assert.False(t, ok4)

	cache.Set("age", []byte("24"))
	_, ok5 := cache.hashmap["age"]
	assert.True(t, ok5)

	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)
}

func TestCacheOverflow(t *testing.T) {
	cache := GetCache(180)
	assert.Equal(t, cache.maxBytes, 180)
	assert.Equal(t, cache.curBytes, 0)
	assert.Equal(t, cache.head.next, &cache.rear)
	assert.Equal(t, cache.rear.prev, &cache.head)
	assert.Equal(t, len(cache.hashmap), 0)

	originItem := &Item{
		key:   "name1",
		value: []byte("jerry"),
	}
	originItemSize := originItem.Size()

	cache.Set("name1", []byte("jerry"))
	assert.Equal(t, cache.curBytes, originItemSize)
	assert.Equal(t, cache.curBytes, 66)
	assert.Equal(t, cache.head.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev, &cache.head)
	assert.Equal(t, len(cache.hashmap), 1)

	cache.Set("name2", []byte("jerry"))
	assert.Equal(t, cache.curBytes, originItemSize*2)
	assert.Equal(t, cache.curBytes, 132)
	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)
	assert.Equal(t, len(cache.hashmap), 2)

	cache.Set("name3", []byte("jerry"))
	assert.Equal(t, cache.curBytes, originItemSize*2)
	assert.Equal(t, cache.curBytes, 132)
	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)
	assert.Equal(t, len(cache.hashmap), 2)
}

func TestCacheGet(t *testing.T) {
	cache := GetCache(180)
	assert.Equal(t, cache.maxBytes, 180)
	assert.Equal(t, cache.curBytes, 0)
	assert.Equal(t, cache.head.next, &cache.rear)
	assert.Equal(t, cache.rear.prev, &cache.head)
	assert.Equal(t, len(cache.hashmap), 0)

	originItem := &Item{
		key:   "name1",
		value: []byte("jerry"),
	}
	originItemSize := originItem.Size()

	name1 := []byte("jerry")
	cache.Set("name1", name1)
	assert.Equal(t, cache.curBytes, originItemSize)
	assert.Equal(t, cache.curBytes, 66)
	assert.Equal(t, cache.head.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev, &cache.head)
	assert.Equal(t, len(cache.hashmap), 1)

	name1, ok1 := cache.Get("name1")
	assert.True(t, ok1)
	assert.Equal(t, string(name1), "jerry")
	name1 = name1[1:]
	assert.Equal(t, string(name1), "erry")
	name1, _ = cache.Get("name1")
	assert.Equal(t, string(name1), "jerry")

	name2, ok2 := cache.Get("Name1")
	assert.False(t, ok2)
	assert.Equal(t, name2, []byte(nil))

	cache.Set("name2", []byte("jerry"))
	assert.Equal(t, cache.curBytes, originItemSize*2)
	assert.Equal(t, cache.curBytes, 132)
	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)
	assert.Equal(t, len(cache.hashmap), 2)

	cache.Set("name3", []byte("jerry"))
	assert.Equal(t, cache.curBytes, originItemSize*2)
	assert.Equal(t, cache.curBytes, 132)
	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)
	assert.Equal(t, len(cache.hashmap), 2)

	// after set `name3`, the `name1` item will be deleted.
	name1, ok1 = cache.Get("name1")
	assert.False(t, ok1)
	assert.Equal(t, name1, []byte(nil))

	name3, ok3 := cache.Get("name2")
	assert.True(t, ok3)
	assert.Equal(t, name3, []byte("jerry"))

	cache.Set("name4", []byte("jerry"))
	assert.Equal(t, cache.curBytes, originItemSize*2)
	assert.Equal(t, cache.curBytes, 132)
	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)
	assert.Equal(t, len(cache.hashmap), 2)

	name4, ok4 := cache.Get("name3")
	assert.False(t, ok4)
	assert.Equal(t, name4, []byte(nil))
}

func TestCacheDelete(t *testing.T) {
	cache := GetCache(180)
	assert.Equal(t, cache.maxBytes, 180)
	assert.Equal(t, cache.curBytes, 0)
	assert.Equal(t, cache.head.next, &cache.rear)
	assert.Equal(t, cache.rear.prev, &cache.head)
	assert.Equal(t, len(cache.hashmap), 0)

	originItem := &Item{
		key:   "name1",
		value: []byte("jerry"),
	}
	originItemSize := originItem.Size()

	cache.Set("name1", []byte("jerry"))
	assert.Equal(t, cache.curBytes, originItemSize)
	assert.Equal(t, cache.curBytes, 66)
	assert.Equal(t, cache.head.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev, &cache.head)
	assert.Equal(t, len(cache.hashmap), 1)

	cache.Set("name2", []byte("jerry"))
	assert.Equal(t, cache.curBytes, originItemSize*2)
	assert.Equal(t, cache.curBytes, 132)
	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)
	assert.Equal(t, len(cache.hashmap), 2)

	cache.Set("name3", []byte("jerry"))
	assert.Equal(t, cache.curBytes, originItemSize*2)
	assert.Equal(t, cache.curBytes, 132)
	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)
	assert.Equal(t, len(cache.hashmap), 2)

	cache.Delete("name1")
	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)
	assert.Equal(t, len(cache.hashmap), 2)

	cache.Delete("name2")
	assert.Equal(t, cache.head.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev, &cache.head)
	assert.Equal(t, len(cache.hashmap), 1)

	cache.Delete("name3")
	assert.Equal(t, cache.head.next, &cache.rear)
	assert.Equal(t, cache.rear.prev, &cache.head)
	assert.Equal(t, len(cache.hashmap), 0)

	assert.Equal(t, cache.curBytes, 0)
}

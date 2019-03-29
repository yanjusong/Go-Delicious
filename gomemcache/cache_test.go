package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheSet(t *testing.T) {
	cache := GetCache(180)
	assert.Equal(t, 180, cache.maxBytes)
	assert.Equal(t, 0, cache.curBytes)
	assert.Equal(t, cache.head.next, &cache.rear)
	assert.Equal(t, cache.rear.prev, &cache.head)
	assert.Equal(t, 0, len(cache.hashmap))

	jerryItem := &Item{
		key:   "name",
		value: []byte("jerry"),
	}
	jerryItemSize := jerryItem.Size()

	cache.Set("name", []byte("jerry"))
	assert.Equal(t, jerryItemSize, cache.curBytes)
	assert.Equal(t, 65, cache.curBytes)
	assert.Equal(t, cache.head.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev, &cache.head)
	item1, ok1 := cache.hashmap["name"]
	assert.True(t, ok1)
	assert.Equal(t, "jerry", string(item1.value))
	assert.Equal(t, 1, len(cache.hashmap))

	cache.Set("name", []byte("jerry"))
	assert.Equal(t, jerryItemSize, cache.curBytes)
	assert.Equal(t, 65, cache.curBytes)
	assert.Equal(t, cache.head.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev, &cache.head)
	item2, ok2 := cache.hashmap["name"]
	assert.True(t, ok2)
	assert.Equal(t, "jerry", string(item2.value))
	assert.Equal(t, 1, len(cache.hashmap))

	cache.Set("name", []byte("jerr"))
	assert.Equal(t, jerryItemSize-1, cache.curBytes)
	assert.Equal(t, 64, cache.curBytes)
	assert.Equal(t, cache.head.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev, &cache.head)
	item3, ok3 := cache.hashmap["name"]
	assert.True(t, ok3)
	assert.Equal(t, "jerr", string(item3.value))
	assert.Equal(t, 1, len(cache.hashmap))

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
	assert.Equal(t, 180, cache.maxBytes)
	assert.Equal(t, 0, cache.curBytes)
	assert.Equal(t, cache.head.next, &cache.rear)
	assert.Equal(t, cache.rear.prev, &cache.head)
	assert.Equal(t, 0, len(cache.hashmap))

	originItem := &Item{
		key:   "name1",
		value: []byte("jerry"),
	}
	originItemSize := originItem.Size()

	cache.Set("name1", []byte("jerry"))
	assert.Equal(t, cache.curBytes, originItemSize)
	assert.Equal(t, 66, cache.curBytes)
	assert.Equal(t, cache.head.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev, &cache.head)
	assert.Equal(t, 1, len(cache.hashmap))

	cache.Set("name2", []byte("jerry"))
	assert.Equal(t, originItemSize*2, cache.curBytes)
	assert.Equal(t, 132, cache.curBytes)
	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)
	assert.Equal(t, 2, len(cache.hashmap))

	cache.Set("name3", []byte("jerry"))
	assert.Equal(t, originItemSize*2, cache.curBytes)
	assert.Equal(t, 132, cache.curBytes)
	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)
	assert.Equal(t, 2, len(cache.hashmap))
}

func TestCacheGet(t *testing.T) {
	cache := GetCache(180)
	assert.Equal(t, 180, cache.maxBytes)
	assert.Equal(t, 0, cache.curBytes)
	assert.Equal(t, cache.head.next, &cache.rear)
	assert.Equal(t, cache.rear.prev, &cache.head)
	assert.Equal(t, 0, len(cache.hashmap))

	originItem := &Item{
		key:   "name1",
		value: []byte("jerry"),
	}
	originItemSize := originItem.Size()

	name1 := []byte("jerry")
	cache.Set("name1", name1)
	assert.Equal(t, originItemSize, cache.curBytes)
	assert.Equal(t, 66, cache.curBytes)
	assert.Equal(t, cache.head.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev, &cache.head)
	assert.Equal(t, 1, len(cache.hashmap))

	name1, ok1 := cache.Get("name1")
	assert.True(t, ok1)
	assert.Equal(t, "jerry", string(name1))
	name1 = name1[1:]
	assert.Equal(t, "erry", string(name1))
	name1, _ = cache.Get("name1")
	assert.Equal(t, "jerry", string(name1))

	name2, ok2 := cache.Get("Name1")
	assert.False(t, ok2)
	assert.Equal(t, []byte(nil), name2)

	cache.Set("name2", []byte("jerry"))
	assert.Equal(t, originItemSize*2, cache.curBytes)
	assert.Equal(t, 132, cache.curBytes)
	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)
	assert.Equal(t, 2, len(cache.hashmap))

	cache.Set("name3", []byte("jerry"))
	assert.Equal(t, originItemSize*2, cache.curBytes)
	assert.Equal(t, 132, cache.curBytes)
	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)
	assert.Equal(t, 2, len(cache.hashmap))

	// after set `name3`, the `name1` item will be deleted.
	name1, ok1 = cache.Get("name1")
	assert.False(t, ok1)
	assert.Equal(t, []byte(nil), name1)

	name3, ok3 := cache.Get("name2")
	assert.True(t, ok3)
	assert.Equal(t, []byte("jerry"), name3)

	cache.Set("name4", []byte("jerry"))
	assert.Equal(t, originItemSize*2, cache.curBytes)
	assert.Equal(t, 132, cache.curBytes)
	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)
	assert.Equal(t, 2, len(cache.hashmap))

	name4, ok4 := cache.Get("name3")
	assert.False(t, ok4)
	assert.Equal(t, []byte(nil), name4)
}

func TestCacheDelete(t *testing.T) {
	cache := GetCache(180)
	assert.Equal(t, 180, cache.maxBytes)
	assert.Equal(t, 0, cache.curBytes)
	assert.Equal(t, cache.head.next, &cache.rear)
	assert.Equal(t, cache.rear.prev, &cache.head)
	assert.Equal(t, 0, len(cache.hashmap))

	originItem := &Item{
		key:   "name1",
		value: []byte("jerry"),
	}
	originItemSize := originItem.Size()

	cache.Set("name1", []byte("jerry"))
	assert.Equal(t, originItemSize, cache.curBytes)
	assert.Equal(t, 66, cache.curBytes)
	assert.Equal(t, cache.head.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev, &cache.head)
	assert.Equal(t, 1, len(cache.hashmap))

	cache.Set("name2", []byte("jerry"))
	assert.Equal(t, originItemSize*2, cache.curBytes)
	assert.Equal(t, 132, cache.curBytes)
	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)
	assert.Equal(t, 2, len(cache.hashmap))

	cache.Set("name3", []byte("jerry"))
	assert.Equal(t, originItemSize*2, cache.curBytes)
	assert.Equal(t, 132, cache.curBytes)
	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)
	assert.Equal(t, 2, len(cache.hashmap))

	cache.Delete("name1")
	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)
	assert.Equal(t, 2, len(cache.hashmap))

	cache.Delete("name2")
	assert.Equal(t, cache.head.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev, &cache.head)
	assert.Equal(t, 1, len(cache.hashmap))

	cache.Delete("name3")
	assert.Equal(t, cache.head.next, &cache.rear)
	assert.Equal(t, cache.rear.prev, &cache.head)
	assert.Equal(t, 0, len(cache.hashmap))

	assert.Equal(t, 0, cache.curBytes)
}

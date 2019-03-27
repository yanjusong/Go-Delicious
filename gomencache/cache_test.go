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

	cache.Set("name", []byte("jerry"))
	assert.Equal(t, cache.curBytes, jerryItemSize)
	assert.Equal(t, cache.curBytes, 65)
	assert.Equal(t, cache.head.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev, &cache.head)
	item2, ok2 := cache.hashmap["name"]
	assert.True(t, ok2)
	assert.Equal(t, string(item2.value), "jerry")

	cache.Set("name", []byte("jerr"))
	assert.Equal(t, cache.curBytes, jerryItemSize-1)
	assert.Equal(t, cache.curBytes, 64)
	assert.Equal(t, cache.head.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev, &cache.head)
	item3, ok3 := cache.hashmap["name"]
	assert.True(t, ok3)
	assert.Equal(t, string(item3.value), "jerr")

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

	cache.Set("name2", []byte("jerry"))
	assert.Equal(t, cache.curBytes, originItemSize*2)
	assert.Equal(t, cache.curBytes, 132)
	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)

	cache.Set("name3", []byte("jerry"))
	assert.Equal(t, cache.curBytes, originItemSize*2)
	assert.Equal(t, cache.curBytes, 132)
	assert.Equal(t, cache.head.next.next.next, &cache.rear)
	assert.Equal(t, cache.rear.prev.prev.prev, &cache.head)
}

func TestCacheGet(t *testing.T) {
	cache := GetCache(180)
	assert.Equal(t, cache.maxBytes, 180)
	assert.Equal(t, cache.curBytes, 0)
	assert.Equal(t, cache.head.next, &cache.rear)
	assert.Equal(t, cache.rear.prev, &cache.head)

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

	name1Get, ok1 := cache.Get("name1")
	assert.True(t, ok1)
	assert.Equal(t, string(name1Get), "jerry")
	name1 = name1[1:]
	assert.Equal(t, string(name1), "erry")
	name1Get, _ = cache.Get("name1")
	assert.Equal(t, string(name1Get), "jerry")

	_, ok2 := cache.Get("Name1")
	assert.False(t, ok2)
}

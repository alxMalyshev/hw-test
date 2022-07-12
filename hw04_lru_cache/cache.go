package hw04lrucache

import (
	"github.com/alxMalyshev/hw-test/hw04_lru_cache/list"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	// Cache // Remove me after realization.
	capacity int
	queue    list.List
	items    map[Key]*list.ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) *lruCache {
	return &lruCache{
		capacity: capacity,
		queue:    list.NewList(),
		items:    make(map[Key]*list.ListItem, capacity),
	}

}

func (l *lruCache) Set(key Key, value interface{}) bool {
	if item, ok := l.items[key]; ok {
		l.queue.MoveToFront(value)
		item.Value = value
		return true
	} else {
		
	}

	return true
}
func (l *lruCache) Get(key Key) bool {return}
func (l *lruCache) Clear() {return}

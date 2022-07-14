package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	// Cache    // Remove me after realization.
	capacity int
	queue    List
	items    map[Key]*ListItem
}

// type cacheItem struct {
// 	key   Key
// 	value interface{}
// }

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	if item, ok := l.items[key]; ok {
		l.queue.MoveToFront(item)
		item.Value = value
		return true
	}
	if l.queue.Len() == l.capacity {
		l.Clear()
	}
	newListItem := l.queue.PushFront(value)
	l.items[key] = newListItem

	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := l.items[key]; ok {
		l.queue.MoveToFront(item)
		return item.Value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	if item := l.queue.Back(); item != nil {
		l.queue.Remove(item)
		for k, v := range l.items {
			if item == v {
				delete(l.items, k)
			}
		}
	}
}

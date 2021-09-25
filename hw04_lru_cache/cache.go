package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (lru *lruCache) Set(key Key, value interface{}) bool {
	lru.mu.Lock()
	defer lru.mu.Unlock()
	item, ok := lru.items[key]
	cItem := cacheItem{key: key, value: value}

	if ok && item != nil {
		item.Value = cItem
		lru.queue.MoveToFront(item)
	} else {
		lru.queue.PushFront(cItem)
	}
	lru.items[key] = lru.queue.Front()

	if lru.queue.Len() > lru.capacity {
		lru.deleteLast()
	}

	return ok
}

func (lru *lruCache) Get(key Key) (interface{}, bool) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	item, ok := lru.items[key]
	if !ok {
		return nil, false
	}

	lru.queue.MoveToFront(item)
	cItem := item.Value.(cacheItem)

	return cItem.value, true
}

func (lru *lruCache) Clear() {
	lru.mu.Lock()
	defer lru.mu.Unlock()
	lru.items = make(map[Key]*ListItem)
	lru.queue = NewList()
}

func (lru *lruCache) deleteLast() {
	lastItem := lru.queue.Back()
	lru.queue.Remove(lastItem)
	cItem := lastItem.Value.(cacheItem)
	delete(lru.items, cItem.key)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

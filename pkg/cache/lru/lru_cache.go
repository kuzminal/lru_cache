package lru

import "container/list"

type Item struct {
	Key   string
	Value string
}

type LRU struct {
	capacity int
	items    map[string]*list.Element
	queue    *list.List
}

func (L *LRU) Add(key, value string) bool {
	if element, exists := L.items[key]; exists == true {
		L.queue.MoveToFront(element)
		return false
	}

	if L.queue.Len() == L.capacity {
		L.removeLastElement()
	}

	item := &Item{
		Key:   key,
		Value: value,
	}

	element := L.queue.PushFront(item)
	L.items[item.Key] = element

	return true
}

func (L *LRU) Get(key string) (value string, ok bool) {
	element, exists := L.items[key]
	if !exists {
		return "", false
	}
	L.queue.MoveToFront(element)
	return element.Value.(*Item).Value, true
}

func (L *LRU) Remove(key string) (ok bool) {
	element, exists := L.items[key]
	if exists {
		L.queue.Remove(element)
		return true
	} else {
		return false
	}
}

func (L *LRU) removeLastElement() {
	if element := L.queue.Back(); element != nil {
		item := L.queue.Remove(element).(*Item)
		delete(L.items, item.Key)
	}
}

func NewLRUCache(n int) Cache {
	return &LRU{
		capacity: n,
		items:    make(map[string]*list.Element),
		queue:    list.New(),
	}
}

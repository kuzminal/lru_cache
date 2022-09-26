package main

import (
	"LRU_cache/pkg/cache/lru"
	"log"
)

func main() {
	cache := lru.NewLRUCache(2)
	log.Println(cache.Add("key1", "value1"))
	log.Println(cache.Add("key2", "value2"))
	log.Println(cache.Get("key1"))
	log.Println(cache.Add("key3", "value3"))
	log.Println(cache.Get("key2"))
}

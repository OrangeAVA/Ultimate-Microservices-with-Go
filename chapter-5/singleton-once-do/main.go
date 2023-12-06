package main

import "sync"

type cache struct{}

var cacheSingleton *cache

var once sync.Once

func main() {
	GetCache()
	GetCache()
	GetCache()
}

func GetCache() *cache {
	once.Do(func() {
		println("Creating singleton object")
		cacheSingleton = &cache{}
	})
	return cacheSingleton
}

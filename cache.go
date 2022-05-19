package cache

import (
	"time"
)

type cacheDatum struct {
	value          string
	expirationTime time.Time
}

type Cache struct {
	data map[string]cacheDatum
}

func NewCache() Cache {
	return Cache{}
}

func (cache *Cache) Get(key string) (string, bool) {
	cachedValue, isValueExist := cache.data[key]

	if !isValueExist {
		return "", false
	}

	if !cachedValue.expirationTime.After(time.Now()) {
		delete(cache.data, key)
		return "", false
	}

	return cachedValue.value, true
}

func (cache *Cache) Put(key, value string) {
	cache.data[key] = cacheDatum{value: value}
}

func (cache *Cache) Keys() []string {
	cacheSize := len(cache.data)
	keys := make([]string, cacheSize)
	i := 0
	for key, datum := range cache.data {
		if datum.expirationTime.After(time.Now()) {
			keys[i] = key
			i++
		} else {
			delete(cache.data, key)
		}
	}
	numberOfExpiredValues := cacheSize - i

	return keys[0 : cacheSize-numberOfExpiredValues]
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	if deadline.After(time.Now()) {
		cache.data[key] = cacheDatum{value: value, expirationTime: deadline}
	}
}

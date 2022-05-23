package cache

import (
	"fmt"
	"time"
)

type Cache struct {
	data map[string]string
}

func NewCache() Cache {
	return Cache{data: map[string]string{}}
}

func (c *Cache) Get(key string) (string, bool) {
	value, ok := c.data[key]
	return value, ok
}

func (c *Cache) Put(key, value string) {
	c.data[key] = value
}

func (c *Cache) Keys() []string {
	keys := []string{}
	for key := range c.data {
		keys = append(keys, key)
	}
	return keys
}

func removeFromMap(dict map[string]string, key string) func() {
	return func() {
		delete(dict, key)
	}
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	duration := deadline.Sub(time.Now())
	fmt.Println("Duration:", duration)
	time.AfterFunc(duration, removeFromMap(c.data, key))
	c.data[key] = value
}

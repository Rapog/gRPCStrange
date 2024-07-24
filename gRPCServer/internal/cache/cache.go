package cache

import (
	"sync"
)

type Cache[T any] interface {
	Set(key string, value T)
	Get(key string) T
	Delete(key string)
	Has(key string) bool
}

type cache[T any] struct {
	mu   sync.Mutex
	data map[string]T
}

func (c *cache[T]) Set(key string, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value

	//log.Println("Set: ", key)
}

func (c *cache[T]) Get(key string) T {
	c.mu.Lock()
	defer c.mu.Unlock()

	//log.Println("Get: ", key)
	return c.data[key]
}

func (c *cache[T]) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	//log.Println("Delete: ", key)
	delete(c.data, key)
}

func (c *cache[T]) Has(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.data[key]
	//log.Println("Has: ", key, " ok: ", ok)
	return ok
}

func New[T any](cap int) Cache[T] {
	//log.Println("New: ", cap)
	return &cache[T]{
		mu:   sync.Mutex{},
		data: make(map[string]T, cap),
	}
}

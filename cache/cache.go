package cache

import (
	"sync"
)

// Interface - реализуйте этот интерфейс
type Interface interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

// Не меняйте названия структуры и название метода создания экземпляра Cache, иначе не будут проходить тесты

type Cache struct {
	// TODO: ваш код
	data map[string]string
	mu   sync.RWMutex
}

// NewCache создаёт и возвращает новый экземпляр Cache.
func NewCache() Interface {
	// TODO: ваш код
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Set(k, v string) {
	// TODO implement me
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[k] = v
}

func (c *Cache) Get(k string) (v string, ok bool) {
	// TODO implement me
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[k]
	return val, ok
}

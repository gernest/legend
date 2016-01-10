package config

import (
	"sync"
)

var cfg = &config{}

type config struct {
	mu   sync.RWMutex
	data map[interface{}]interface{}
}

func (c *config) set(key, val interface{}) {
	c.mu.Lock()
	c.data[key] = val
	c.mu.Unlock()
}

func (c *config) get(key interface{}) interface{} {
	c.mu.RLock()
	data := c.data[key]
	c.mu.RUnlock()
	return data
}

func Set(key interface{}, value interface{}) {
	cfg.set(key, value)
}

func Get(key interface{}) interface{} {
	return cfg.get(key)
}

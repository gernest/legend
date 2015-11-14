package config

import (
	"sync"
)

var cfg = &config{}

type config struct {
	mu   sync.Mutex
	data map[interface{}]interface{}
}

func (c *config) set(key, val interface{}) {
	c.mu.Lock()
	c.data[key] = val
	c.mu.Unlock()
}

func (c *config) get(key interface{}) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.data[key]
}

func Set(key interface{}, value interface{}) {
	cfg.set(key, value)
}

func Get(key interface{}) interface{} {
	return cfg.get(key)
}

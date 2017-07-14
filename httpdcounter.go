package httpdcounter

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu     sync.Mutex
	values map[string]int64
}

func New() *Counter {
	return &Counter{
		values: make(map[string]int64),
	}
}

func (c *Counter) Get(key string) int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.values[key]
}

func (c *Counter) Set(key string) int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[key]++
	return c.values[key]
}

func (c *Counter) String() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	var s string
	for k, v := range c.values {
		s += fmt.Sprintf("[%s] %d\n", k, v)
	}
	return s
}

func (c *Counter) GetMap() *map[string]int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return &c.values
}

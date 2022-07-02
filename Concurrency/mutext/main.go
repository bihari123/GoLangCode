package main

import "sync"

type Container struct {
	mu  sync.Mutex
	counters map[string]int
}

func (c *Container) increment(name string) {
	c.mu.Lock
}

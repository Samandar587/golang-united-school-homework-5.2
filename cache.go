package main

import (
	"time"
)

type Cache struct {
	kv       map[string]string
	ev       map[string]string
	deadline map[string]time.Time
	dead     map[string]bool
}

func NewCache() Cache {
	c := Cache{
		kv:       make(map[string]string),
		ev:       make(map[string]string),
		deadline: make(map[string]time.Time),
		dead:     make(map[string]bool),
	}
	return c
}

func (c Cache) Get(key string) (string, bool) {

	if val, ok := c.kv[key]; ok {
		return val, true
	}

	if val, ok := c.ev[key]; ok {
		if c.deadline[key].Before(time.Now()) {
			c.dead[key] = true
			return "", false
		}
		return val, true
	}

	return "", false
}

func (c Cache) Put(key, value string) {
	c.kv[key] = value
}

func (c Cache) Keys() []string {
	keys := make([]string, len(c.kv))

	i := 0
	for k := range c.kv {
		keys[i] = k
		i++
	}

	for k := range c.ev {
		if c.dead[k] == false {
			keys = append(keys, k)
		}
	}

	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.ev[key] = value
	c.deadline[key] = deadline
}

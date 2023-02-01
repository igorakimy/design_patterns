// Интерфейс стратегии.

package main

type EvictionAlgo interface {
	evict(c *Cache)
}

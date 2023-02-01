// Конкретная стратегия.

package main

import "fmt"

type Lru struct {
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicted by lru strategy")
}

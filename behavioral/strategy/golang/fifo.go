// Конкретная стратегия.

package main

import "fmt"

type Fifo struct {
}

func (f *Fifo) evict(c *Cache) {
	fmt.Println("Evicted by fifo strategy")
}

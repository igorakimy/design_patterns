// Конкретная стратегия.

package main

import "fmt"

type Lfu struct {
}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicted by lfu strategy")
}

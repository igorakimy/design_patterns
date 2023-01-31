package main

import "fmt"

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	fmt.Scanln()
}

// Creating single instance now.
// Single instance already created.
// Single instance already created.
// Single instance already created.

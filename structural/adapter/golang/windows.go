// Неизвестный сервис.

package main

import "fmt"

type Windows struct {
}

func (w *Windows) InsertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

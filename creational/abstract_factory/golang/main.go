/*
Абстрактная фабрика — это порождающий паттерн проектирования,
который решает проблему создания целых семейств связанных продуктов,
без указания конкретных классов продуктов.

Абстрактная фабрика задаёт интерфейс создания всех доступных типов
продуктов, а каждая конкретная реализация фабрики порождает продукты
одной из вариаций. Клиентский код вызывает методы фабрики для
получения продуктов, вместо самостоятельного создания с помощью
оператора new. При этом фабрика сама следит за тем, чтобы создать
продукт нужной вариации.
*/
package main

import "fmt"

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

// Logo: nike
// Size: 14
// Logo: nike
// Size: 14
// Logo: adidas
// Size: 14
// Logo: adidas
// Size: 14

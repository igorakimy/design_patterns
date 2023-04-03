// Клиентский код
package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()

	if err := walletFacade.addMoneyToWallet("abc", 1234, 10); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	if err := walletFacade.deductMoneyFromWallet("abc", 1234, 5); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}

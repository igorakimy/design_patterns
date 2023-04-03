// Часть сложной подсистемы
package main

import "fmt"

type Account struct {
	name string
}

func newAccount(accountID string) *Account {
	return &Account{
		name: accountID,
	}
}

func (a *Account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

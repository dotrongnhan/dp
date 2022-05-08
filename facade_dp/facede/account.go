package facede

import "fmt"

type Account struct {
	Name string
}

func NewAccount(accountName string) *Account {
	return &Account{
		Name: accountName,
	}
}

func (a *Account) CheckAccount(accountName string) error {
	if a.Name != accountName {
		return fmt.Errorf("account's name is incorrect")
	}
	fmt.Println("Account verified")
	return nil
}

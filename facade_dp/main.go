package main

import (
	"facade/facede"
	"fmt"
	"log"
)

func main() {
	fmt.Println("=====")
	wallFacade := facede.NewWalletFacade("IICAOD", 100001)
	fmt.Println("=========>")
	err := wallFacade.AddMoneyToWallet("IICAOD", 100001, 100000)
	if err != nil {
		log.Fatalf("ERROR: %s\n", err.Error())
	}
	fmt.Println("=========>")
	err = wallFacade.DeductMoneyFromWallet("IICAD", 100001, 3000)
	if err != nil {
		log.Fatalf("ERROR: %s\n", err.Error())
	}
}

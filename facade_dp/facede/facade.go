package facede

import "fmt"

type WalletFacade struct {
	Account      *Account
	Wallet       *Wallet
	SecurityCode *SecurityCode
	Notification *Notification
}

func NewWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	return &WalletFacade{
		Account:      NewAccount(accountID),
		SecurityCode: NewSecurityCode(code),
		Notification: &Notification{},
		Wallet:       NewWallet(),
	}
}

func (w *WalletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.Account.CheckAccount(accountID)
	if err != nil {
		return err
	}

	err = w.SecurityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}

	w.Wallet.creditBalance(amount)
	w.Notification.sendWalletCreditNotification()
	return nil
}

func (w *WalletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money to wallet")
	err := w.Account.CheckAccount(accountID)
	if err != nil {
		return err
	}

	err = w.SecurityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}

	err = w.Wallet.debitBalance(amount)
	if err != nil {
		return err
	}

	w.Notification.sendWalletDebitNotification()
	return nil
}

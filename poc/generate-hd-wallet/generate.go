package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"log"
)

const mnemonic = "barely snack visit march rather area struggle budget aisle pilot muscle surprise" /* local test pk */

func newWalletFromMnemonic() *hdwallet.Wallet {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}
	return wallet
}

func Generate(index int) accounts.Account {
	wallet := newWalletFromMnemonic()
	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", index))
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}
	return account
}

func GetKeypair(account accounts.Account) (string, string) {
	wallet := newWalletFromMnemonic()
	pvk, err := wallet.PrivateKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}
	pub, err := wallet.PublicKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}
	return pvk, pub
}

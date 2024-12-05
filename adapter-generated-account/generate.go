package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"log"
	"os"
)

const fallbackMnemonic = "barely snack visit march rather area struggle budget aisle pilot muscle surprise" /* fallback pk */

func newWalletFromMnemonic() *hdwallet.Wallet {
	mnemonic, existed := os.LookupEnv("MNEMONIC")
	if !existed {
		log.Println("Mnemonic environment variable not set, Use fallback mnemonic")
		mnemonic = fallbackMnemonic
	}

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}
	return wallet
}

func Generate(change int, index int) accounts.Account {
	wallet := newWalletFromMnemonic()
	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/%d'/0/%d", change, index))
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

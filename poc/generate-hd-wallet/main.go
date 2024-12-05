package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"runtime"
	"sync"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {

	generatedAccount := map[string]*accounts.Account{}
	generatedPrivateKey := map[string]string{}

	fmt.Println("Version", runtime.Version())
	fmt.Println("NumCPU", runtime.NumCPU())
	//fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(60))

	batch := 3_000_000
	iterate := 20
	var wg sync.WaitGroup
	wg.Add(iterate)
	for i := 0; i < iterate; i++ {
		go execute(i, i*batch, batch, generatedAccount, generatedPrivateKey, &wg)
	}
	wg.Wait()
	fmt.Println("Done!")
}

func execute(iterate int, from int, to int, generatedAccount map[string]*accounts.Account, generatedPrivateKey map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := from; i < to; i++ {

		/* Generate account */
		account := Generate(i)
		/* Check is an account existed & put */
		_, accountExisted := generatedAccount[account.Address.String()]
		if accountExisted {
			panic(fmt.Sprintf("[%d] Account %s already existed", iterate, account.Address.String()))
		}
		generatedAccount[account.Address.String()] = &account

		/* Generate a private key & public key */
		pvk, pub := GetKeypair(account)
		_, privateKeyExisted := generatedPrivateKey[pvk]
		if privateKeyExisted {
			panic(fmt.Sprintf("[%d] Private key %s already existed", iterate, pvk))
		}
		generatedPrivateKey[pvk] = pvk

		fmt.Printf("[%d] Account: %v, Pvk: %v, Pub: %v\n", iterate, account, pvk, pub)
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

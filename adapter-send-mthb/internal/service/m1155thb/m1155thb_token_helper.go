package m1155thb

import (
	"adapter-send/internal/model"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func ExecuteSendM1155BTHBToken(client *ethclient.Client, payload *model.Payload1155, contractAddress string) {
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Error on view chain id... %v", err)
	}
	fmt.Printf("Chain ID : %v \n", chainId)

	privateKey, err := crypto.HexToECDSA(payload.IssuerPrivateKey)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatalf("Error on initial new transaction with chain id: %v (%v)", err, chainId)
	}

	contract := initialContract(client, contractAddress)
	to := common.HexToAddress(payload.To)
	opt := &bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		//Value:  big.NewInt(0),
	}

	/* Batch Transfer */
	//amounts := []*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(1)}
	//ids := []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2), big.NewInt(3)}
	//tx, err := contract.SafeBatchTransferFrom(opt, auth.From, to, ids, amounts, []byte{})
	//if err != nil {
	//	log.Fatalf("Error on execute transaction: %v ", err)
	//}

	/* Single Transfer */
	amount := &payload.Amount
	id := &payload.Id
	tx, err := contract.SafeTransferFrom(opt, auth.From, to, id, amount, []byte{})
	if err != nil {
		log.Fatalf("Error on execute transaction: %v ", err)
	}

	fmt.Printf("Success execute store pending transaction [%v] (%v => %v)\n", tx.Hash(), payload.From, payload.To)
}

func ExecuteReadM1155BTHBToken(client *ethclient.Client, payload *model.Payload1155, contractAddress string) {
	contract := initialContract(client, contractAddress)

	from := common.HexToAddress(payload.From)
	callOpts := bind.CallOpts{
		Pending: false,
		Context: nil,
	}

	/* Batch Balance Of */
	froms := []common.Address{from, from, from, from}
	ids := []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2), big.NewInt(3)}
	balances, err := contract.BalanceOfBatch(&callOpts, froms, ids)
	if err != nil {
		log.Fatalf("Error on call balance: %v ", err)
	}
	log.Printf("Balance [0] : %v \n", balances[0])
	log.Printf("Balance [1] : %v \n", balances[1])
	log.Printf("Balance [2] : %v \n", balances[2])
	log.Printf("Balance [3] : %v \n", balances[3])

	/* Single Balance Of */
	//balance, err := contract.BalanceOf(&callOpts, from, big.NewInt(0))
	//if err != nil {
	//	log.Fatalf("Error on call balance: %v ", err)
	//}
	//log.Printf("Balance [0] : %v \n", balance)

}

func initialContract(client *ethclient.Client, contractAddress string) *Internal {
	address := common.HexToAddress(contractAddress)
	contract, err := NewInternal(address, client)
	if err != nil {
		panic("Unable to create contract instance")
	}
	return contract
}

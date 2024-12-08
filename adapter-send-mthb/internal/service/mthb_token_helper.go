package internal

import (
	"adapter-send/internal/model"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func ExecuteSendMBTHToken(client *ethclient.Client, payload *model.Payload, contractAddress string) {
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
	amount := &payload.Amount
	opt := &bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		//Value:  big.NewInt(0),
	}
	//opt.GasPrice, _ = client.SuggestGasPrice(context.Background())
	tx, err := contract.Transfer(opt, to, amount)
	if err != nil {
		log.Fatalf("Error on execute transaction: %v ", err)
	}
	fmt.Printf("Success execute store pending transaction [%v] (%v => %v)\n", tx.Hash(), payload.From, payload.To)
}

func initialContract(client *ethclient.Client, contractAddress string) *Internal {
	address := common.HexToAddress(contractAddress)
	contract, err := NewInternal(address, client)
	if err != nil {
		panic("Unable to create contract instance")
	}
	return contract
}

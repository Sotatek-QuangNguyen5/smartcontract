package main

import (
    
	"crypto/ecdsa"
	"fmt"
	"log"
    "context"
    "math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	contract "geth-server/mycontract" // the contract packages you generated with abigen
)

func main() {
    // initialize eth client
    client, err := ethclient.Dial("HTTP://127.0.0.1:8545")
    if err != nil {
        log.Fatal(err)
    }

    // load the signing key
    privateKey, err := crypto.HexToECDSA("13a8f11b976f5469ba7f50de3f6670a4db4d034fb8f6c33a4afa754185fed94b")
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    // estimate gas price
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(5e18)     // in wei
    auth.GasLimit = uint64(3000000) // in units
    auth.GasPrice = gasPrice

    // your deployed smart contract address
    address := common.HexToAddress("0x667441Be0AE98fd4bC527712C6b2Fc6F6700fcFf")

    // get an instance of your smart contract
    instance, err := contract.NewMyContract(address, client)
    if err != nil {
        log.Fatal(err)
    }

    // calling your smart contract method
    tx, err := instance.List(nil)
    if err != nil {
        log.Fatal(err)
    }

    // transaction hash
    // fmt.Printf("tx sent: %s", tx.Hash().Hex()) 
    fmt.Println(tx)
}
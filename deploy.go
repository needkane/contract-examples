package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/needkane/contract-examples/mytoken"
)

const key = `
{
    "address": "ce30fbd14fa07ca27e296e32927a4d2467ee2f54",
    "crypto": {
        "cipher": "aes-128-ctr",
        "ciphertext": "19248a8bc8d2dedb6efc67e263c3233636a17d9322d7fea4e4dbf08569fcb8a0",
        "cipherparams": {
            "iv": "102a57406016e7d929f885200eb435b9"
        },
        "kdf": "scrypt",
        "kdfparams": {
            "dklen": 32,
            "n": 262144,
            "p": 1,
            "r": 8,
            "salt": "578911f486d1f21e5a8fe9e6b6252dc8fae1c66046e6fb9d6e76d2db435815cc"
        },
        "mac": "946250b9c58408cca3363ed3298b4f4cd563629d2329c941badc4faff4f24272"
    },
    "id": "8cb35aea-be5e-4b2c-83a9-9fa2a6a82f4f",
    "version": 3
}
`

func main() {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	conn, err := ethclient.Dial("http://10.253.4.251:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), "123456")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// Deploy a new awesome contract for the binding demo
	//	address, tx, token, err := mytoken.DeployContract(auth, conn, big.NewInt(9651), "Contracts in Go!!!", 0, "Go!")
	address, tx, token, err := mytoken.DeployContract(auth, conn, big.NewInt(9651), "Contracts in Go!!!", 0, "Go!")
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	startTime := time.Now()
	fmt.Printf("TX start @:%s", time.Now())
	ctx := context.Background()
	addressAfterMined, err := bind.WaitDeployed(ctx, conn, tx)
	if err != nil {
		log.Fatalf("failed to deploy contact when mining :%v", err)
	}
	fmt.Printf("tx mining take time:%s\n", time.Now().Sub(startTime))
	if bytes.Compare(address.Bytes(), addressAfterMined.Bytes()) != 0 {
		log.Fatalf("mined address :%s,before mined address:%s", addressAfterMined, address)
	}
	name, err := token.Name(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Pending name:", name)
}

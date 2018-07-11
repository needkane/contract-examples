package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	// Create an IPC based RPC connection to a remote node and instantiate a contract binding
	conn, err := ethclient.Dial("http://10.253.4.251:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	token, err := mytoken.NewMyToken(common.HexToAddress("f71462113f8539177bf6e0c2c6a555c03ddee94d"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	toAddress := common.HexToAddress("8c1b2e9e838e2bf510ec7ff49cc607b718ce8401")
	val, _ := token.BalanceOf(nil, toAddress)
	fmt.Printf("before transfer :%s\n", val)
	// Create an authorized transactor and spend 1 unicorn
	auth, err := bind.NewTransactor(strings.NewReader(key), "123456")
	auth.GasLimit = 66317
	fmt.Println(auth)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	//fmt.Println(auth)
	tx, err := token.Transfer(auth, toAddress, 387)
	fmt.Printf("======61-=========%x", tx.Hash())
	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
	}
	ctx := context.Background()
	receipt, err := bind.WaitMined(ctx, conn, tx)
	if err != nil {
		log.Fatalf("tx mining error:%v\n", err)
	}
	val, _ = token.BalanceOf(nil, toAddress)
	fmt.Printf("after transfere:%s\n", val)
	fmt.Printf("tx is :%s\n", tx)
	fmt.Printf("receipt is :%s\n", receipt)
}

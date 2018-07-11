package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nkbai/go-callcontract-example/mytoken"
)

func main() {
	// Create an IPC based RPC connection to a remote node and instantiate a contract binding
	conn, err := ethclient.Dial("http://10.253.4.251:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	token, err := mytoken.NewMyToken(common.HexToAddress("0xf71462113f8539177bf6e0c2c6a555c03ddee94d"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	contractName, err := token.Name(nil)
	if err != nil {
		log.Fatalf("query name err:%v", err)
	}
	fmt.Printf("MyToken Name is:%s\n", contractName)
	balance, err := token.BalanceOf(nil, common.HexToAddress("0x8c1b2e9e838e2bf510ec7ff49cc607b718ce8401"))
	if err != nil {
		log.Fatalf("query balance error:%v", err)
	}
	fmt.Printf("0x8c1b2e9e838e2bf510ec7ff49cc607b718ce8401's balance is %s\n", balance)
}

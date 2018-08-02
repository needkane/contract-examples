package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privStr := "88e5e86e6a15e318b9ac433a67ad0b0532bf1a22b6af7ac94481d866be3e7759"
	privkey, err := crypto.ToECDSA(common.Hex2Bytes(privStr))
	if err != nil {
		log.Fatal(" ToECDSA failed: ", err)
	}

	toStr := "0x145f7376df0e16d94ae1d0d5fa8f00a96d99a6e3"
	valueI := 11
	tokenStr := "0x610033b6dd5a08004e46f2097ca09b693d744118"
	expireUnix := 1633205071
	sequenceI := 1

	toAddr := common.HexToAddress(toStr).Bytes()
	value := Int64ToBytes(valueI)
	tokenAddr := common.HexToAddress(tokenStr).Bytes()
	expireTime := Int64ToBytes(expireUnix)
	sequenceId := Int64ToBytes(sequenceI)

	bytez := crypto.Keccak256([]byte("ERC20"), toAddr, value, tokenAddr, expireTime, sequenceId)
	fmt.Printf("Keccak256:%x\n", bytez)
	sig, err := crypto.Sign(bytez, privkey)
	if err != nil {
		log.Fatal("Sign failed: ", err)
	}
	fmt.Printf("signature: 0x%x", sig)
}

func Int64ToBytes(i int) []byte {
	var buf = make([]byte, 32)
	PutUint256(buf, uint64(i))
	return buf
}

func PutUint256(b []byte, v uint64) {
	_ = b[31] // early bounds check to guarantee safety of writes below
	b[0] = byte(v >> 248)
	b[1] = byte(v >> 240)
	b[2] = byte(v >> 232)
	b[3] = byte(v >> 224)
	b[4] = byte(v >> 216)
	b[5] = byte(v >> 208)
	b[6] = byte(v >> 200)
	b[7] = byte(v >> 192)
	b[8] = byte(v >> 184)
	b[9] = byte(v >> 176)
	b[10] = byte(v >> 168)
	b[11] = byte(v >> 160)
	b[12] = byte(v >> 152)
	b[13] = byte(v >> 144)
	b[14] = byte(v >> 136)
	b[15] = byte(v >> 128)
	b[16] = byte(v >> 120)
	b[17] = byte(v >> 112)
	b[18] = byte(v >> 104)
	b[19] = byte(v >> 96)
	b[20] = byte(v >> 88)
	b[21] = byte(v >> 80)
	b[22] = byte(v >> 72)
	b[23] = byte(v >> 64)
	b[24] = byte(v >> 56)
	b[25] = byte(v >> 48)
	b[26] = byte(v >> 40)
	b[27] = byte(v >> 32)
	b[28] = byte(v >> 24)
	b[29] = byte(v >> 16)
	b[30] = byte(v >> 8)
	b[31] = byte(v)
}

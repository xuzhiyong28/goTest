package go_merkle

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func makeDividendAccounts() []DividendAccount {
	accounts := make([]DividendAccount, 0)
	accounts = append(accounts, NewDividendAccount(common.HexToAddress("0x3F2E3dA83cbA1C2e128CAedeE7CeF6a8bF33C8dd"), "10000"))
	accounts = append(accounts, NewDividendAccount(common.HexToAddress("0xeC7705Efe4A40B07fD8293F0b9Ba8E60554EBa0E"), "20000"))
	accounts = append(accounts, NewDividendAccount(common.HexToAddress("0x16B0eC57e4308D2eD96Da643aB45973397Ceb980"), "30000"))
	return accounts
}

func TestRootHash(t *testing.T) {
	accounts := makeDividendAccounts()
	tree, err := GetAccountTree(accounts)
	if err == nil {
		fmt.Println(fmt.Sprintf("rootHash = %v", "0x"+hex.EncodeToString(tree.Root.Hash)))
	}
}

func TestAccountProof(t *testing.T) {
	accounts := makeDividendAccounts()
	proof, index, err := GetAccountProof(accounts, common.HexToAddress("0xeC7705Efe4A40B07fD8293F0b9Ba8E60554EBa0E"))
	if err == nil {
		fmt.Println(fmt.Sprintf("proof = %v, index = %v", "0x"+hex.EncodeToString(proof), index))
	}
}


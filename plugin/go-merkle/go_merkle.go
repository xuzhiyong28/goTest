package go_merkle

import (
	"bytes"
	"fmt"
	"github.com/cbergoon/merkletree"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"math/big"
	"sort"
)

type Address [20]byte

func (a Address) Bytes() []byte { return a[:] }

type DividendAccount struct {
	User      common.Address `json:"user"`
	FeeAmount string         `json:"feeAmount"`
}

func NewDividendAccount(user common.Address, fee string) DividendAccount {
	return DividendAccount{
		User:      user,
		FeeAmount: fee,
	}
}

func (da *DividendAccount) String() string {
	if da == nil {
		return "nil-DividendAccount"
	}
	return fmt.Sprintf("DividendAccount{%s %v}", da.User, da.FeeAmount)
}

// 实现 merkle相关的函数
func (da DividendAccount) CalculateHash() ([]byte, error) {
	fee, _ := big.NewInt(0).SetString(da.FeeAmount, 10)
	divAccountHash := crypto.Keccak256(appendBytes32(da.User.Bytes(), fee.Bytes()))
	return divAccountHash, nil
}

func (da DividendAccount) Equals(other merkletree.Content) (bool, error) {
	if bytes.Compare(da.User.Bytes(), common.Address{}.Bytes()) == 0 && bytes.Compare(other.(DividendAccount).User.Bytes(), common.Address{}.Bytes()) == 0 {
		return true, nil
	}
	return bytes.Equal(da.User.Bytes(), other.(DividendAccount).User.Bytes()), nil
}

// 排序
func SortDividendAccountByAddress(dividendAccounts []DividendAccount) []DividendAccount {
	sort.Slice(dividendAccounts, func(i, j int) bool {
		return bytes.Compare(dividendAccounts[i].User.Bytes(), dividendAccounts[j].User.Bytes()) < 0
	})
	return dividendAccounts
}

func GetAccountTree(dividendAccounts []DividendAccount) (*merkletree.MerkleTree, error) {
	// 1. 按照账号排序
	dividendAccounts = SortDividendAccountByAddress(dividendAccounts)
	// 2. 构造排序后的list
	var list []merkletree.Content
	for i := 0; i < len(dividendAccounts); i++ {
		list = append(list, dividendAccounts[i])
	}
	tree, err := merkletree.NewTreeWithHashStrategy(list, sha3.NewLegacyKeccak256)
	if err != nil {
		return nil, err
	}
	return tree, nil
}

func GetAccountProof(dividendAccounts []DividendAccount, userAddress common.Address) ([]byte, uint64, error) {
	dividendAccounts = SortDividendAccountByAddress(dividendAccounts)
	var list []merkletree.Content
	var account DividendAccount
	index := uint64(0)
	for i := 0; i < len(dividendAccounts); i++ {
		list = append(list, dividendAccounts[i])
		if bytes.Equal(dividendAccounts[i].User.Bytes(), userAddress.Bytes()) {
			account = dividendAccounts[i]
			index = uint64(i)
		}
	}
	tree, err := merkletree.NewTreeWithHashStrategy(list, sha3.NewLegacyKeccak256)
	if err != nil {
		return nil, 0, err
	}
	branchArray, _, err := tree.GetMerklePath(account)
	proof := appendBytes32(branchArray...)
	return proof, index, err
}

func appendBytes32(data ...[]byte) []byte {
	var result []byte
	for _, v := range data {
		paddedV, err := convertTo32(v)
		if err == nil {
			result = append(result, paddedV[:]...)
		}
	}
	return result
}

func convertTo32(input []byte) (output [32]byte, err error) {
	l := len(input)
	if l > 32 || l == 0 {
		return
	}
	copy(output[32-l:], input[:])
	return
}

package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

func main() {
	key, _ := crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	sig, err := crypto.Sign(crypto.Keccak256([]byte("foo")), key)
	if err == nil {
		R, S, V := decodeSignature(sig)
		fmt.Println(R)
		fmt.Println(S)
		fmt.Println(V)
	}
}

func decodeSignature(sig []byte) (r, s, v *big.Int) {
	if len(sig) != crypto.SignatureLength {
		panic(fmt.Sprintf("wrong size for signature: got %d, want %d", len(sig), crypto.SignatureLength))
	}
	r = new(big.Int).SetBytes(sig[:32])
	s = new(big.Int).SetBytes(sig[32:64])
	v = new(big.Int).SetBytes([]byte{sig[64] + 27})
	return r, s, v
}

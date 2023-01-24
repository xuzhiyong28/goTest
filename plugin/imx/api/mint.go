package api

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/crypto"
)

type TokenJson struct {
	Id        string `json:"id"`
	Blueprint string `json:"blueprint"`
}

type UsersParam struct {
	Address string      `json:"ether_key"`
	Tokens  []TokenJson `json:"tokens"`
}

type SignaturePayloadObject struct {
	ContractAddress string       `json:"contract_address"`
	Users           []UsersParam `json:"users"`
	AuthSignature   string       `json:"auth_signature"`
}

func (spo *SignaturePayloadObject) ToJsonStr() string {
	str, _ := json.Marshal(spo)
	return string(str)
}

// 签名
func (spo *SignaturePayloadObject) ImxSignature(privateKey string) (string, error) {
	privateKeyEcdsa, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	msg := crypto.Keccak256([]byte(spo.ToJsonStr()))
	sig, err := crypto.Sign(msg, privateKeyEcdsa)
	if err == nil {
		R, S, V := DecodeSignatureStr(sig)
		str := R + S + V
		return str, nil
	}
	return "", nil
}

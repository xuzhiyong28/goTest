package api

import (
	"encoding/json"
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

	return "", nil
}

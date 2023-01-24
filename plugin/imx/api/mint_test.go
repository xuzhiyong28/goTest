package api

import (
	"fmt"
	"testing"
)

func TestToJson(t *testing.T) {
	tokenJsons := make([]TokenJson, 0)
	tokenJsons = append(tokenJsons, TokenJson{
		Id:        "1",
		Blueprint: "1",
	}, TokenJson{
		Id:        "2",
		Blueprint: "2",
	})
	usersParams := make([]UsersParam, 0)
	usersParams = append(usersParams, UsersParam{
		Address: "0x3F2E3dA83cbA1C2e128CAedeE7CeF6a8bF33C8dd",
		Tokens:  tokenJsons,
	})

	signaturePayloadObject := SignaturePayloadObject{
		ContractAddress: "0x8d439da20b756f26ac70714296063d14a5ff2507",
		Users:           usersParams,
		AuthSignature:   "",
	}

	fmt.Println(signaturePayloadObject.ToJsonStr())
	signature, err := signaturePayloadObject.ImxSignature("f4ac35bdca06836ca46dce8e84490c2676a1416384bc99ecc388785ac7dc40ca")
	if err == nil {
		fmt.Println(signature)
	}
}

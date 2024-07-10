package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
	"github.com/immutable/imx-core-sdk-golang/imx/signers/ethereum"
	"math/big"
	"testing"
)

var (
	contractAddress = "0xb9ce71061bffbdd7268b5c3292c7153ae987e00b"
	userAddress     = "0x3F2E3dA83cbA1C2e128CAedeE7CeF6a8bF33C8dd"
	privateKey      = "f4ac35bdca06836ca46dce8e84490c2676a1416384bc99ecc388785ac7dc40ca"
)

func initializeSDK() *api.APIClient {
	configuration := api.NewConfiguration()
	return api.NewAPIClient(configuration)
}

// 获取资产
func TestAssetDemo1(t *testing.T) {
	client := initializeSDK()
	assetsRequest := api.ApiListAssetsRequest{}
	assetsRequest = assetsRequest.User(contractAddress)
	assetsRequest = assetsRequest.Collection(userAddress)
	assetsRequest = assetsRequest.PageSize(100)
	assetsResponse, httpResponse, err := client.AssetsApi.ListAssetsExecute(assetsRequest)
	if err == nil && httpResponse.StatusCode == 200 {
		printJson(assetsResponse)
	}
}

func TestAssetDemo2(t *testing.T) {
	client := initializeSDK()
	tokenId := "146447"
	assetRequest := client.AssetsApi.GetAsset(context.Background(), contractAddress, tokenId)
	assetsResponse, httpResponse, err := client.AssetsApi.GetAssetExecute(assetRequest)
	if err == nil && httpResponse.StatusCode == 200 {
		printJson(&assetsResponse)
	}
}

func TestMintDemo1(t *testing.T) {
	signer, _ := ethereum.NewSigner(privateKey, new(big.Int).SetInt64(5))
	client := initializeSDK()
	apiMintTokensRequest := client.MintsApi.MintTokens(context.Background())
	token1 := api.NewMintTokenDataV2("3")
	token1.SetBlueprint("3")
	token2 := api.NewMintTokenDataV2("4")
	token2.SetBlueprint("4")
	mintRequest := api.NewMintRequest("", contractAddress, []api.MintUser{
		{
			Tokens: []api.MintTokenDataV2{
				*token1,
				*token2,
			},
			User: userAddress,
		},
	})
	message, _ := mintRequest.MarshalJSON()
	messageHash := crypto.Keccak256Hash(message)
	signMessage, err := signer.SignMessage(messageHash.String())
	if err == nil {
		fmt.Println(signMessage)
	}
	mintRequest.SetAuthSignature(hexutil.Encode(signMessage))
	apiMintTokensRequest = apiMintTokensRequest.MintTokensRequestV2([]api.MintRequest{*mintRequest})
	response, httpResponse, err := client.MintsApi.MintTokensExecute(apiMintTokensRequest)
	if err == nil && httpResponse.StatusCode == 200 {
		fmt.Println(response)
	}

}

func printJson(obj interface{}) {
	str, _ := json.Marshal(&obj)
	fmt.Println(string(str))
}

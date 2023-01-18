package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
	"testing"
)

var (
	contractAddress = "0xb9ce71061bffbdd7268b5c3292c7153ae987e00b"
	userAddress     = "0x3F2E3dA83cbA1C2e128CAedeE7CeF6a8bF33C8dd"
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

func TestBalanceDemo1(t *testing.T) {

}

func printJson(obj interface{}) {
	str, _ := json.Marshal(&obj)
	fmt.Println(string(str))
}

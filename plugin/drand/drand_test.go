package drand

import (
	"bytes"
	"context"
	"fmt"
	"github.com/drand/drand/chain"
	"github.com/drand/drand/client"
	"github.com/drand/drand/client/http"
	"log"
	"testing"
	"time"
)

var urls = []string{
	"https://api.drand.sh",
	"https://drand.cloudflare.com",
}

// 2020-07-22 23:17:30 - 1595431050
var chainInfoJSON = `
{
  "public_key": "868f005eb8e6e4ca0a47c8a77ceaa5309a47978a7c71bc5cce96366b5d7a569937c529eeda66c7293784a9402801af31",
  "period": 30,
  "genesis_time": 1595431050,
  "hash": "8990e7a9aaed2ffed73dbd7092123d6f289930540d7651336225dc172e51b2ce",
  "groupHash": "176f93498eac9ca337150b46d21dd58673ea4e3581185f869672e59fa4cb390a"
}
`

// 2020-05-26 06:19:35 - 1590445175
var chaininfoJSON_Test = `
{
  "public_key": "922a2e93828ff83345bae533f5172669a26c02dc76d6bf59c80892e12ab1455c229211886f35bb56af6d5bea981024df",
  "period": 25,
  "genesis_time": 1590445175,
  "hash": "84b2234fb34e835dccd048255d7ad3194b81af7d978c3bf157e3469592ae4e02",
  "groupHash": "4dd408e5fdff9323c76a9b6f087ba8fdc5a6da907bd9217d9d10f2287d081957"
}
`

func TestDemo0(t *testing.T) {
	setup1 := 885004
	setup2 := 885005
	drandChain, _ := chain.InfoFromJSON(bytes.NewReader([]byte(chainInfoJSON)))

	client, err := client.New(
		client.From(http.ForURLs(urls, drandChain.Hash())...),
		client.WithChainHash(drandChain.Hash()),
	)
	if err != nil {
		log.Fatal(err)
	}
	resp1, _ := client.Get(context.TODO(), uint64(setup1))
	fmt.Println(
		fmt.Sprintf("round : %v\n signature: %v\n randomness: %v\n", resp1.Round(), resp1.Signature(), resp1.Randomness()),
	)

	resp2, _ := client.Get(context.TODO(), uint64(setup2))
	fmt.Println(
		fmt.Sprintf("round : %v\n signature: %v\n randomness: %v\n", resp2.Round(), resp2.Signature(), resp2.Randomness()),
	)
	err = chain.VerifyBeacon(drandChain.PublicKey, &chain.Beacon{
		PreviousSig: resp1.Signature(),
		Round:       resp2.Round(),
		Signature:   resp2.Signature(),
	})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("verifyBeacon successful")
	}

}

func TestDemo1(t *testing.T) {

	drandChain, _ := chain.InfoFromJSON(bytes.NewReader([]byte(chainInfoJSON)))

	client, err := client.New(
		client.From(http.ForURLs(urls, drandChain.Hash())...),
		client.WithChainHash(drandChain.Hash()),
	)
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= 100; i++ {
		resp, err := client.Get(context.TODO(), uint64(i))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(
			fmt.Sprintf("round : %v\n signature: %v\n randomness: %v\n", resp.Round(), resp.Signature(), resp.Randomness()),
		)
		time.Sleep(30 * time.Second)
	}
}

func TestMaxBeaconRoundForEpoch(t *testing.T) {
	epoch := MaxBeaconRoundForEpoch(1, 1621981175, 30)
	fmt.Println(fmt.Sprintf("epoch : %v", epoch))
}

func MaxBeaconRoundForEpoch(filEpoch uint64, filGenTime uint64, filRoundTime uint64) uint64 {
	drandChain, _ := chain.InfoFromJSON(bytes.NewReader([]byte(chainInfoJSON)))
	latestTs := filEpoch * filRoundTime + filGenTime - 30
	dround := (latestTs - uint64(drandChain.GenesisTime)) / uint64(drandChain.Period.Seconds())
	return dround
}

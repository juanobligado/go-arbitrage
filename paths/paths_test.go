package paths

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/joho/godotenv"
	"github.com/juanobligado/go-arbitrage/tokenmetadata"
	"github.com/juanobligado/go-arbitrage/uniswap_pair"
	"github.com/juanobligado/go-arbitrage/utils"
)

func init(){
	godotenv.Load()
}

func TestReadPaths(m *testing.T){
	dict := tokenmetadata.Restore("../data/tokenAddressMap.json")
	pairMetadata := ReadPaths("../data/uni_sushi_paths.json",dict)
	fmt.Print(pairMetadata)
}

func TestCalculateAllPaths(m *testing.T){
	dict := tokenmetadata.Restore("../data/tokenAddressMap.json")
	pairMetadata := ReadPaths("../data/uni_sushi_paths.json",dict)
	prices :=  uniswap_pair.RestorePrices("../data/pair_balances.json")

	pathResults := CalculateAll(pairMetadata,prices)
	Save(pathResults,"../data/path_calc_result.json")
}

func TestCalculatePaths(m *testing.T){
	dict := tokenmetadata.Restore("../data/tokenAddressMap.json")
	pairMetadata := ReadPaths("../data/uni_sushi_paths.json",dict)
	prices :=  uniswap_pair.RestorePrices("../data/pair_balances.json")
	result := pairMetadata[0].Calculate(prices)
	fmt.Print(result)
}




func TestCalculate(m *testing.T){
	ars := tokenmetadata.Token{
		Symbol: "ARS",
		Address: "$ARS",
	}
	usd := tokenmetadata.Token{
		Symbol: "USD",
		Address: "$USD",
	}
	arsusdGovt := uniswap_pair.PairBalance{
		Info: uniswap_pair.PairMetadata{
			Address: "arsusdGovt",
			T0: ars,
			T1: usd,
		},
		T0Balance: utils.BigInt{Int: *big.NewInt(100*1000)},
		T1Balance: utils.BigInt{Int: *big.NewInt(1000)},
	}
	arsusdFree := uniswap_pair.PairBalance{
		Info: uniswap_pair.PairMetadata{
			Address: "arsusdFree",
			T0: ars,
			T1: usd,
		},
		T0Balance: utils.BigInt{Int: *big.NewInt(300*1000)},
		T1Balance: utils.BigInt{Int: *big.NewInt(1000)},
	}
	prices := make(uniswap_pair.PairBalances)
	prices[arsusdFree.Info.Address] = arsusdFree
	prices[arsusdGovt.Info.Address] = arsusdGovt


	path :=  Path{
							Index: 0,
							Trivial: false,
							Tokens: []string{"ARS" ,"USD"},
							PathItems:  []PathItem{
								{ PairAddress: arsusdFree.Info.Address, DstToken: ars.Address},
								{ PairAddress: arsusdGovt.Info.Address, DstToken: usd.Address},
							},
				}
	
	result := path.Calculate(prices)
	fmt.Print(result.Ratio)
	
}


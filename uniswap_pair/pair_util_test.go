package uniswap_pair

import (
	"fmt"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/juanobligado/go-arbitrage/token_metadata"
)
func TestCreatePairMetadata(m *testing.T) {

	client, err := ethclient.Dial(getInfuraKey())
    if err != nil {
        log.Fatal(err)
    }
	dict := token_metadata.CreateFromList("../data/tokens.json")
	metadata := ReadPairMetadata("0x05f04f112a286c4c551897fb19ed2300272656c8",dict.Dictionary,client)
	fmt.Print(metadata)

}

func TestGenerateMetadataPairFile(m *testing.T){
	client, err := ethclient.Dial(getInfuraKey())
    if err != nil {
        log.Fatal(err)
    }
	
	dict := token_metadata.CreateFromList("../data/tokens.json")
	addresses:= token_metadata.ReadDistinctPairs("../data/uni_sushi_paths.json")
	//addresses := []string{"0x05f04f112a286c4c551897fb19ed2300272656c8"} 
	GeneratePairMetadataFile("../data/pair_metadata.json",addresses,dict.Dictionary,client)

}

func getInfuraKey() string {
	return "https://mainnet.infura.io/v3/ce50544cce4f4619b8a32afe1a8b06e4"
}



func TestReadPairPrice(m *testing.T){
	client, err := ethclient.Dial(getInfuraKey())
    if err != nil {
        log.Fatal(err)
    }
	
	dict := token_metadata.CreateFromList("../data/tokens.json")
	addresses := []string{"0x05f04f112a286c4c551897fb19ed2300272656c8"} 
	balances,_ := ReadPairPrices(addresses,dict.Dictionary,client)
	fmt.Println(balances)
}

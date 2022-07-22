package uniswap_pair

import (
	"fmt"
	"log"
	"testing"

	"github.com/joho/godotenv"
	M "github.com/juanobligado/go-arbitrage/tokenmetadata"
)

func init(){
	godotenv.Load()
}

func TestCreatePairMetadata(m *testing.T) {

	client, err := CreateInfuraConnection()
    if err != nil {
        log.Fatal(err)
    }	
	dict := M.CreateFromList("../data/tokens.json")
	metadata := ReadPairMetadata("0x05f04f112a286c4c551897fb19ed2300272656c8",dict.Dictionary,client)
	fmt.Print(metadata)

}

func TestGenerateMetadataPairFile(m *testing.T){
	client, err := CreateInfuraConnection()
    if err != nil {
        log.Fatal(err)
    }
	
	dict := M.CreateFromList("../data/tokens.json")
	addresses:= M.ReadDistinctPairs("../data/uni_sushi_paths.json")
	GeneratePairMetadataFile("../data/pair_metadata.json",addresses,dict.Dictionary,client)

}


// Test to generate Price File for all balances
func TestReadAllPrices(m *testing.T){
	pairMetadata := RestorePairMetadata("../data/pair_metadata.json")

	client, err := CreateInfuraConnection()
    if err != nil {
        log.Fatal(err)
    }
	balances,_ := ReadPairPrices(pairMetadata,client)

	err = balances.Save("../data/pair_balances.json")
	if err != nil {
        log.Fatal(err)
    }
}

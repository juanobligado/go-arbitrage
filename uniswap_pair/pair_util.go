package uniswap_pair

import (
	"math/big"
	"strings"

	"encoding/json"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/juanobligado/go-arbitrage/token_metadata"
)

type PairMetadata struct {
	Address string `json:"address"`
	T0 token_metadata.Token `json:"t0"`
	T1 token_metadata.Token `json:"t1"`
}

type PairBalance struct{
	T0Balance big.Int
	T1Balance big.Int
}

func ReadPairMetadata(address string,  dictionary token_metadata.TokenByAddress, client  *ethclient.Client ) PairMetadata {


	metadata :=  PairMetadata{ Address: address }

	// Read Pair Metadata From Blockchain
	pairInstance,err := NewUniswapPairCaller(common.HexToAddress(address),client)
	if(err!=nil){
		return metadata
	}
	t0_address,err := pairInstance.Token0(nil)
	if(err!=nil){
		return metadata
	}
	t1_address,err := pairInstance.Token1(nil)
	if(err!=nil){
		return metadata
	}
	
	str_t0_address := strings.ToLower(t0_address.Hex())  
	str_t1_address := strings.ToLower(t1_address.Hex())
	
	metadata.T0 = dictionary[str_t0_address]
	metadata.T1 = dictionary[str_t1_address]
	
	return metadata
}

// Creates a File with Extended Metadata for Token Pairs
func  GeneratePairMetadataFile(filename string,addresses []string,dictionary token_metadata.TokenByAddress, client  *ethclient.Client) error {

	pairMetadataMap := make(map[string]PairMetadata)

	for i:=0;i<len(addresses);i++{
		address := addresses[i]
		if  _ ,ok := pairMetadataMap[address]; !ok  {
			pairMetadataMap[address] = ReadPairMetadata(addresses[i],dictionary,client)
		}
	}
	addressToMetadataFile, err := json.MarshalIndent(pairMetadataMap, "", " ") 
	if(err!= nil){
		return err
	}
	return ioutil.WriteFile(filename, addressToMetadataFile, 0644)
}


func ReadPairPrices(strAddresses []string,  dictionary token_metadata.TokenByAddress, client  *ethclient.Client ) (map[string]PairBalance ,error) {


	uniswapViewContractAddress :=  "0x416355755f32b2710ce38725ed0fa102ce7d07e6"
	balances := make(map[string]PairBalance)
	addresses := make([]common.Address,len(strAddresses))
	for i:=0;i<len(strAddresses);i++{
		addresses[i] = common.HexToAddress(strAddresses[i])
	}
	// Read Pair Metadata From Blockchain
	priceProxyInstance,err := NewUniswapViewCaller(common.HexToAddress(uniswapViewContractAddress),client)
	if(err!= nil){
		return balances,err
	}
	result,err := priceProxyInstance.ViewPair(nil,addresses)
	for i:=0;i<len(strAddresses);i++{
		item := PairBalance{}
		item.T0Balance = *result[i*2]
		item.T1Balance = *result[i*2+1]
		balances[strAddresses[i]] = item
	}
	return balances,nil
}
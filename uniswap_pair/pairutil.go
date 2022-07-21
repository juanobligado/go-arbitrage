package uniswap_pair

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/juanobligado/go-arbitrage/config"
	T "github.com/juanobligado/go-arbitrage/tokenmetadata"
)

type PairMetadata struct {
	Address string `json:"address"`
	T0 T.Token `json:"t0"`
	T1 T.Token `json:"t1"`
}

type PairBalance struct{
	Info PairMetadata 
	T0Balance big.Int
	T1Balance big.Int
}

type PairBalances map[string]PairBalance

func ReadPairMetadata(address string,  dictionary T.TokenByAddress, client  *ethclient.Client ) PairMetadata {


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
func  GeneratePairMetadataFile(filename string,addresses []string,dictionary T.TokenByAddress, client  *ethclient.Client) error {

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

func  RestorePairMetadata(filename string) map[string]PairMetadata {

	instance := make(map[string]PairMetadata)

	filedata,_ := ioutil.ReadFile(filename)
	err := json.Unmarshal(filedata, &instance) 
	if(err!= nil){
		return nil
	}
	return instance
}

type emtpyArgError struct{error }
func ReadPairPrices(pairs map[string]PairMetadata , client  *ethclient.Client ) (PairBalances ,error) {



	uniswapViewContractAddress := config.New().GoArbitrageConfig.PriceProxyContractAddress 
	balances := make(PairBalances)
	if(len(pairs) == 0){
		return balances , emtpyArgError{} 
	}
	addresses := make([]common.Address,len(pairs))

	i:=0
	for  k,_ := range  pairs{
		addresses[i] = common.HexToAddress(k)
		i++
	}
	// Read Pair Metadata From Blockchain
	priceProxyInstance,err := NewUniswapViewCaller(common.HexToAddress(uniswapViewContractAddress),client)
	if(err!= nil){
		return balances,err
	}
	result,err := priceProxyInstance.ViewPair(nil,addresses)
	if(err!=nil){
		return balances,err
	}
	for i:=0;i<len(addresses);i++{
		pairAddress :=  strings.ToLower(addresses[i].String()) 
		item := PairBalance{}
		item.Info = pairs[pairAddress]
		item.T0Balance = *result[i*2]
		item.T1Balance = *result[i*2+1]
		balances[pairAddress] = item
	}
	return balances,nil
}

func (b *PairBalances) Save(filename string ) error{
	marshalledData, err := json.MarshalIndent( (*b), "", " ") 
	if(err!= nil){
		return err
	}
	return ioutil.WriteFile(filename, marshalledData, 0644)
}
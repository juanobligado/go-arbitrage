package paths

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/juanobligado/go-arbitrage/tokenmetadata"
	"github.com/juanobligado/go-arbitrage/uniswap_pair"
	"github.com/juanobligado/go-arbitrage/utils"
)
 

type PathItem struct{
	PairAddress string
	DstToken string
	
}
type Path struct{
	Index int
	Tokens []string
	PathItems []PathItem
	Trivial bool	
}

type RawPathsInterface interface {}
// Read Paths
// restores token metadata from original .json file 
func  ReadPaths(filename string, dictionary *tokenmetadata.TokenDictionary) []Path {

	file, _ := ioutil.ReadFile(filename)
	rawData := []RawPathsInterface{}
 	err := json.Unmarshal([]byte(file), &rawData)
	if(err!=nil){
		return nil;
	}

	paths := make([]Path,0)
 	for iPath := 0; iPath < len(rawData); iPath++ {
		// need to create path summary
		rawPath := rawData[iPath].([]interface{})
		newItems := ReadPathItems(rawPath,iPath,dictionary)
		paths = append(paths, newItems...)

	}		
	// fills metadata dictionary by token address 
	return  paths
}

//Reads Possible path from Path Node
func ReadPathItems(rawPath []interface{}, iPath int, dictionary *tokenmetadata.TokenDictionary) []Path{

	pathItemCount := len(rawPath)
	pathCache := make([]Path,0)
	pathPairs := make(map[string]bool , 0)
	pathCacheItem := Path{Index: iPath}
	pathCacheItem.Tokens = append(pathCacheItem.Tokens,"WETH")
	

	for iPathItem := 0; iPathItem < pathItemCount ; iPathItem++{
		// Reading path item 
		
		rawPathPairs := rawPath[iPathItem].([]interface{})[1].([]interface{})
		dstToken := rawPath[iPathItem].([]interface{})[0].(string)
		dstTokenMeta := dictionary.Dictionary[dstToken]
		pathCacheItem.Tokens = append(pathCacheItem.Tokens, dstTokenMeta.Symbol)
		// TODO: Handle multiple Internal Paths
		for _,rawPathPair := range rawPathPairs{

			strPair :=fmt.Sprint(rawPathPair)
			pathPairs[strPair] = true;	
		}
		newPathItem := PathItem {
			PairAddress: fmt.Sprint(rawPathPairs[0]),
			DstToken: dstToken,
		}

		pathCacheItem.PathItems = append(pathCacheItem.PathItems,newPathItem)
		
	}
	// Just skip Trivial Paths since they dont provide arb opportunities i.e A -> B , B -> A is always 1
	if(len(pathPairs) > 1){
		pathCache = append(pathCache, pathCacheItem)
	}
	
	return pathCache
}

type PathCalcResult struct{
	Nums  []utils.BigInt
	Den   []utils.BigInt
	NTotal utils.BigInt
	DTotal utils.BigInt
	Ratio float64
}



func (p *Path) Calculate( balances uniswap_pair.PairBalances) PathCalcResult {
	result := PathCalcResult{}
	result.DTotal.SetInt64(1)
	result.NTotal.SetInt64(1)

	for _,item := range p.PathItems{
		qNum := utils.BigInt{}
		qDen := utils.BigInt{}
		//todo :calculate multiple paths or check the one with Biggest volatility		
		pairBalance := balances[item.PairAddress]


		if item.DstToken == pairBalance.Info.T0.Address {
			qNum = pairBalance.T0Balance
			qDen = pairBalance.T1Balance
		}else if item.DstToken == pairBalance.Info.T1.Address{
			qNum = pairBalance.T1Balance
			qDen = pairBalance.T0Balance
		}
		result.Nums = append(result.Nums,qNum)
		result.Den  = append(result.Den,qDen)
		result.NTotal.Mul( &result.NTotal.Int, &qNum.Int )
		result.DTotal.Mul( &result.DTotal.Int, &qDen.Int )
	}
	if(result.DTotal.BitLen()!= 0){
		den :=new(big.Float).SetInt(&result.DTotal.Int)
		num :=new(big.Float).SetInt(&result.NTotal.Int)
		result.Ratio,_ = (*new(big.Float).Quo( num,den)).Float64()	
	}
	
	return result
}

func CalculateAll(paths []Path ,balances uniswap_pair.PairBalances) []PathCalcResult{
	result := make([]PathCalcResult,0)
	for _,item := range paths{
		result = append(result, item.Calculate(balances))
	}
	return result
}

func  Save( b []PathCalcResult,filename string ) error{
	marshalledData, err := json.MarshalIndent( b, "", " ") 
	if(err!= nil){
		return err
	}
	return ioutil.WriteFile(filename, marshalledData, 0644)
}
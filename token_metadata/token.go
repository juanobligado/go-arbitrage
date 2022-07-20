// This script transforms token.json file which
package token_metadata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)
 
type Tokens struct {
	tokens []Token 
}
 
type Token struct {
	Address string `json:"address"`
	Symbol string `json:"symbol"`
	Name string`json:"name"`
	Decimals string`json:"decimals"`
}

// token met
type TokenByAddress map[string]Token

type TokenDictionary struct{
	Dictionary TokenByAddress

}

// restores token metadata from original .json file 
func  CreateFromList(filename string) *TokenDictionary{

	file, _ := ioutil.ReadFile(filename)
	data := []Token{}
 	err := json.Unmarshal([]byte(file), &data)
	if(err!=nil){
		return nil;
	}
	instance  := TokenDictionary{}
	instance.Dictionary = make(TokenByAddress)
	// fills metadata dictionary by token address 
	for i := 0; i < len(data); i++ {
		instance.Dictionary[data[i].Address] = data[i]
	}
	return  &instance
}

// restores token map from file
func  Restore(filename string) *TokenDictionary {

	filedata,_ := ioutil.ReadFile(filename)
	instance  := TokenDictionary{}
	err := json.Unmarshal(filedata, &instance.Dictionary) 
	if(err!= nil){
		return nil
	}
	return &instance
}

// saves token metadata map into file 
func (t *TokenDictionary) Save(filename string) error {

	addressToMetadataFile, err := json.MarshalIndent(t.Dictionary, "", " ") 
	if(err!= nil){
		return err
	}
	return ioutil.WriteFile(filename, addressToMetadataFile, 0644)
}

type RawMetadata interface {

}


// restores token metadata from original .json file 
func  ReadDistinctPairs(filename string) []string {

	file, _ := ioutil.ReadFile(filename)
	rawData := []RawMetadata{}
 	err := json.Unmarshal([]byte(file), &rawData)
	if(err!=nil){
		return nil;
	}
	
	distinctPairs := []string{}
	pairMap := make(map[string]bool)
 	for iPath := 0; iPath < len(rawData); iPath++ {
		// need to create path summary
		path := rawData[iPath].([]interface{})
		for jPathElement := 0; jPathElement < len(path) ; jPathElement++{
			rawPairs := path[jPathElement].([]interface{})[1].([]interface{})
			for _,pairAddressNode := range rawPairs{
				pairAddress := fmt.Sprint(pairAddressNode)
				if(!pairMap[pairAddress]){
					pairMap[pairAddress] = true
					distinctPairs = append(distinctPairs, pairAddress)
			
				}		
			}
		}
	}		
	// fills metadata dictionary by token address 
	return  distinctPairs
}
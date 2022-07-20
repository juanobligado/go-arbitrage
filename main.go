package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	Token "github.com/juanobligado/go-arbitrage/token_metadata"
	Pair "github.com/juanobligado/go-arbitrage/uniswap_pair"
)

func main(){
	  
	dict := Token.CreateFromList("./data/tokens.json")
	dict.Save("./data/tokenAddressMap.json")

	address := common.HexToAddress("0xcc3d1ecef1f9fd25599dbea2755019dc09db3c54")
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/ce50544cce4f4619b8a32afe1a8b06e4")
    if err != nil {
        log.Fatal(err)
    }
	instance,err := Pair.NewUniswapPairCaller(address,client)

	address0,_:= instance.Token0(nil)
	address1,_ := instance.Token1(nil)
    fmt.Println("Pair Addresses:",address0,address1)
}
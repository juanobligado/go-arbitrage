package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	Token "github.com/juanobligado/go-arbitrage/tokenmetadata"
	Pair "github.com/juanobligado/go-arbitrage/uniswap_pair"
)

func init(){
	godotenv.Load()
}

func main(){
	  
	dict := Token.CreateFromList("./data/tokens.json")
	dict.Save("./data/tokenAddressMap.json")

	address := common.HexToAddress("0xcc3d1ecef1f9fd25599dbea2755019dc09db3c54")
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/")
    if err != nil {
        log.Fatal(err)
    }
	instance,err := Pair.NewUniswapPairCaller(address,client)

	address0,_:= instance.Token0(nil)
	address1,_ := instance.Token1(nil)
    fmt.Println("Pair Addresses:",address0,address1)
}
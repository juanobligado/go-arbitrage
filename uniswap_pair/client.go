package uniswap_pair

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/juanobligado/go-arbitrage/config"
)

func CreateInfuraConnection() (*ethclient.Client,error){
	config := config.New()
	infura_url := fmt.Sprintf("https://mainnet.infura.io/v3/%s",config.GoArbitrageConfig.InfuraUrl)   
	return ethclient.Dial(infura_url)
}
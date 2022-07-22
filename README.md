# go-arbitrage
 
# 1. Get Pair Metadata  

# 2 Get Balances for all 

To get balances to all pairs we take metadata file Generated on step 1 and call contract to get balances

```
    pairMetadata := RestorePairMetadata("../data/pair_metadata.json")
	client, err := CreateInfuraConnection()
    if err != nil {
        log.Fatal(err)
    }
	balances,_ := ReadPairPrices(pairMetadata,client)
	balances.Save("../data/pair_balances.json")
```

# 3 Process balances and identify arbitrage opportunities

Here we are going to only detect imbalances, will be skipping tx cost estimation which would be needed to further decide if the arb is economically viable

1. Read All Pair Balances genereted in #2
2. Read Paths
3. For Each path Calculated propagated WETH price (should be 1 if no abitrage) 
4. Return a sorted list
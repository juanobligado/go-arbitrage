# go-arbitrage
 
# 1. Get Pair Metadata  

In order to get extended pair metadata. 
1. First created a method to get unique pair addresses  
2. Created method to iterate over a pair collection and created an extended metadata file with t0 and t1 info

```
func TestGenerateMetadataPairFile(m *testing.T){
	client, err := CreateInfuraConnection()
    if err != nil {
        log.Fatal(err)
    }
	
	dict := M.CreateFromList("../data/tokens.json")
	addresses:= M.ReadDistinctPairs("../data/uni_sushi_paths.json")
	GeneratePairMetadataFile("../data/pair_metadata.json",addresses,dict.Dictionary,client)

}
```


[pair_metadata.json](./data/pair_metadata.json)

# 2 Get Balances for each pair

To get balances to all pairs we take metadata file Generated on step 1 and call *ReadPairPrices* contract to get balances and call Proxy contract

```
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
```
1. We load extended pair metadata from #1
2. Call UniswapView proxy contract to fetch all Balances
3. Post Process output and generate data to feed the next step
3. If finding that a pair has some liquidity issues we flag result with liquidity warning  
4. Save Result in File

[pair_balances.json](./data/pair_balances.json)

# 3 Process balances and identify arbitrage opportunities

Here we are going to only detect imbalances, will be skipping tx cost estimation which would be needed to further decide if the arb is economically viable

1. Read All Pair Balances genereted in #2
2. Read Paths 
3. For Each path we do calculate propagated WETH price (should be 1 if no abitrage) 
4. Write down result list as Json for post processing

[path_calc_result.json](./data/path_calc_result.json)

# 4 TODOS
- Enhance Path processing to create multiple possible paths for the same node 
- Extra filtering and data cleaning
- Include Slippage and Transaction costs into calculation
- Refactoring, always refactoring 
- Performance Tuning, (Matrix Operations?, Paralel computing?)
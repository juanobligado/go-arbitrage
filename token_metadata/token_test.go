package token_metadata

import (
	"testing"
)

func TestExportTokenAddressMap(m *testing.T){

	dict := CreateFromList("../data/tokens.json")
	dict.Save("./data/tokenAddressMap.json")

}

func TestReadDistinctPairs(m *testing.T){
	data:= ReadDistinctPairs("../data/uni_sushi_paths.json")
	m.Log(data)
}
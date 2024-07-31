package main

import (
	"fmt"

	"github.com/dgryski/go-farm"
	"github.com/dgryski/go-spooky"
	"github.com/shawnohare/go-minhash"
)

var (
	h1   = spooky.Hash64
	h2   = farm.Hash64
	size = 100
)

type Item struct {
	hostname string
	mh       *minhash.MinHash
	similar  []*Item
}

var dataset []Item

func main() {

	//dummyHostnames := GetAgentHostnames(10, 10, 1)
	//ShowExampleSimilarities(GetRandomHostnames(20))
	//StringToMinHash("dfgh+ghj")

	//vlHostanmes := GetVaryingLenghtSimpleHostnames(12)
	//slices.Reverse(vlHostanmes)
	//ShowExampleSimilarities(vlHostanmes)

	mockHostnames0 := GetVaryingLenghtHostnames(1000)
	mockHostnames1 := GetVaryingLenghtHostnames(1000)
	// mh0 := mockHostnames0[len(mockHostnames0)-1]
	// mh1 := mockHostnames1[len(mockHostnames1)-1]

	//ShowHeatmapInaccuracy(GetAgentHostnames(5, 100, 1134), GetAgentHostnames(5, 100, 1134))
	ShowHeatmapInaccuracy(mockHostnames0, mockHostnames1)

	// // ShowExampleSimilarities(append(mockHostnames0, mockHostnames1[:]...))

	// ShowHeatmapSimilaritiesNonsymmetric(mockHostnames0, mockHostnames0, "", mh0, mh1)

	//fmt.Println(JacardiSimilarity("this_is_", "thesis"))

	// fg := GetFlags(5)
	// fmt.Println("Example flag sequence:" + strings.Join(fg, " "))
	// fg_new := ModifyFlags(1, fg)
	// fmt.Println("Example modified flag sequence:" + strings.Join(fg_new, " "))
	// fmt.Println(StringToMinHash(strings.Join(fg, " ")).Similarity(StringToMinHash(strings.Join(fg_new, " "))))

	fmt.Println(FlagExperiment(10))

}

// mh := StringToMinHash(strings.Join(fg, " "))
// for i := 0; i < 10; i++ {
// 	fg_new := ModifyFlags(1, fg)
// 	mh_new := StringToMinHash(strings.Join(fg_new, " "))
// 	if len(fg_new) == 0 {
// 		fmt.Print("(Note: no flags) ")
// 	}
// 	fmt.Println(mh.Similarity(mh_new))
// 	fg = fg_new
// 	mh = mh_new
// }

package main

import (
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

	mockHostnames0 := GetVaryingLenghtHostnames(100)
	mockHostnames1 := GetVaryingLenghtHostnames(100)
	// mh0 := mockHostnames0[len(mockHostnames0)-1]
	// mh1 := mockHostnames1[len(mockHostnames1)-1]

	//ShowHeatmapInaccuracy(GetAgentHostnames(5, 100, 1134), GetAgentHostnames(5, 100, 1134))
	//ShowHeatmapInaccuracy(GetAgentHostnames(5, 100, 1), GetAgentHostnames(5, 100, 1))

	ShowHeatmapInaccuracy(mockHostnames0, mockHostnames1)

	// // ShowExampleSimilarities(append(mockHostnames0, mockHostnames1[:]...))

	// ShowHeatmapSimilaritiesNonsymmetric(mockHostnames0, mockHostnames0, "", mh0, mh1)

	//fmt.Println(JacardiSimilarity("this_is_", "thesis"))

	//
	// fmt.Println("Example flag sequence:" + strings.Join(fg, " "))
	// fg_new := ModifyFlags(1, fg)
	// fmt.Println("Example modified flag sequence:" + strings.Join(fg_new, " "))
	// fmt.Println(StringToMinHash(strings.Join(fg, " ")).Similarity(StringToMinHash(strings.Join(fg_new, " "))))

	//fmt.Println(FlagExperiment(10))

	/// THE TIME:
	// N := 10000 * 10000
	// fg := GetFlags(10)

	// a := make([]*minhash.MinHash, N+1)
	// a[0] = StringToMinHash(strings.Join(fg, " "))
	// fg_new := ModifyFlags(1, fg)
	// a[1] = StringToMinHash(strings.Join(fg_new, " "))
	// // 	a[i] = StringToMinHash(strings.Join(fg_new, " "))
	// // for i := 0; i < N+1; i++ {
	// // 	fg_new := ModifyFlags(1, fg)
	// // 	a[i] = StringToMinHash(strings.Join(fg_new, " "))
	// // }

	// start := time.Now()
	// for i := 0; i < N; i++ {
	// 	a[0].Similarity(a[1])
	// }

	// stop := time.Since(start).Seconds()
	// fmt.Println(stop)
	// fmt.Println(float64(stop) / float64(N))

	//THE BAT HOSTNAMES
	ShowSimilarityNetwork(GetRealHostnames())

}

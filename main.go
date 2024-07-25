package main

import (
	"fmt"
	"slices"

	"github.com/dgryski/go-farm"
	"github.com/dgryski/go-spooky"
	"github.com/shawnohare/go-minhash"
)

var (
	h1   = spooky.Hash64
	h2   = farm.Hash64
	size = 5
)

type Item struct {
	hostname string
	mh       *minhash.MinHash
	similar  []*Item
}

var dataset []Item

func main() {

	//dummyHostnames := MakeAgentHostnames(10, 1000)
	ShowExampleSimilarities(GetRandomHostnames(20))

	//ShowHeatmapSimilarities(GetAgentHostnames(5, 100, 1134))

	vlHostanmes := GetVaryingLenghtSimpleHostnames(12)

	slices.Reverse(vlHostanmes)

	ShowExampleSimilarities(vlHostanmes)

	fmt.Println(StringToMinHash("this_is_").Signature())
	fmt.Println(StringToMinHash("thesis").Signature())
}

// What do we want?
// A struct with a:
// * hostnames
// * pointer to its mw (which gives us access to a signature for example)
// * ordered list of similar names and their numbers

//Hold list of these.

// Define the data structures. (with the thing set as nil)
// Fill them in.
//   * generate the data
//   * get the similarities in
// Print them up.

// func getSimilar(index int) {
// 	fmt.Printf("%s: ", hostnames[index])
// 	primary := data[index]
// 	similar := []Item{}

// 	for i := 0; i < len(hostnames); i++ {
// 		if primary.Similarity(data[i]) > 0 {
// 			fmt.Printf("%s (%f), ", hostnames[i], primary.Similarity(data[i]))
// 			similar = append(similar, Item{hostname: hostnames[i], similarity: primary.Similarity(data[i])})
// 		}
// 	}

// }

// func FindHostnameSimilarities() {
// 	for i := 0; i < len(dataset); i++ {
// 		item := dataset[i]
// 		candidates := []*Item{}
// 		for j := 0; j < len(dataset); j++ {
// 			if item.mh.Similarity(dataset[j].mh) > 0 {
// 				candidates = append(candidates, &dataset[j])
// 			}
// 		}

// 	}
// }

// for i := 0; i < len(hostnames); i++ {
// 	fmt.Println("")
// 	getSimilar(i)
// }
// mw := minhash.NewMinHash(h1, h2, size) // handle words set

// for _, w := range dummyHostnames[0] {
// 	mw.Push(w)
// }
// fmt.Println(mw.Signature())

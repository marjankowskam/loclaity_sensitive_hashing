package main

import (
	"math/rand"

	"github.com/shawnohare/go-minhash"
)

const sliceLen = 2

func StringToMinHash(word string) *minhash.MinHash {
	slices := []string{}
	for i := 0; i < len(word)-sliceLen+1; i++ {
		slices = append(slices, word[i:i+sliceLen])
	}

	mw := minhash.NewMinHash(h1, h2, size) // handle words set

	for _, w := range slices {
		mw.Push(w)
	}
	return mw
}

func JacardiSimilarity(word0 string, word1 string) float64 {

	set0 := make(map[string]bool)
	set1 := make(map[string]bool)

	for i := 0; i < len(word0)-sliceLen+1; i++ {
		set0[word0[i:i+sliceLen]] = true

	}
	for i := 0; i < len(word1)-sliceLen+1; i++ {
		set1[word1[i:i+sliceLen]] = true

	}
	intersectionSize := 0
	for key := range set0 {
		if set1[key] {
			intersectionSize++
		}
	}
	unionSize := len(set0) + len(set1) - intersectionSize

	return float64(intersectionSize) / float64(unionSize)
}

type BySimilarity struct {
	data        []Item
	mainElement *minhash.MinHash
}

func (b BySimilarity) Len() int {
	return len(b.data)
}

func (b BySimilarity) Swap(i, j int) {
	b.data[i], b.data[j] = b.data[j], b.data[i]
}

func (b BySimilarity) Less(i, j int) bool {
	return b.mainElement.Similarity(b.data[i].mh) > b.mainElement.Similarity(b.data[j].mh)
}

func GetRandomString(num int) string {
	charset := "abcdefghijklmnopqrstuvwxyz0123456789"
	word := ""
	for i := 0; i < num; i++ {
		word += string(charset[rand.Intn(len(charset))])
	}
	return word
}

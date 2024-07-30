package main

import (
	"math/rand"

	"github.com/shawnohare/go-minhash"
)

func StringToMinHash(word string) *minhash.MinHash {
	slices := []string{}
	for i := 1; i < len(word); i++ {
		slices = append(slices, word[i-1:i+1])
	}

	mw := minhash.NewMinHash(h1, h2, size) // handle words set

	for _, w := range slices {
		mw.Push(w)
	}

	return mw
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

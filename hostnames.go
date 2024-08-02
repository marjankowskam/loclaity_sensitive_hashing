package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetAgentHostnames(len int, num int, increment int) []string {
	baseName := "agentname_that_is_long"
	hostnames := []string{}
	hostInt := 0
	for i := 0; i < num; i++ {
		hostnames = append(hostnames, baseName[:len]+strconv.Itoa(hostInt))
		hostInt += increment
	}
	fmt.Println(hostInt)
	return hostnames
}

// Very naive random hostnames
func GetRandomHostnames(num int) []string {
	hostnames := []string{}
	words := []string{"form", "eventsrdftyguh", "love", "oldrdftyg", "johndfgh", "johnny", "main", "call", "hours", "image"}
	for i := 0; i < num; i++ {
		hostnames = append(hostnames, words[rand.Intn(len(words))]+strconv.Itoa(rand.Intn(1000)))
	}
	fmt.Printf("Set of dummy hostnames: %s\n", strings.Join(hostnames, ", "))
	return hostnames
}

func GetVaryingLenghtHostnames(num int) []string {
	baseName := GetRandomString(num)

	hostnames := []string{}
	for i := 0; i < num; i++ {
		hostnames = append(hostnames, baseName[:(i+1)])
	}

	return hostnames
}

func GetRealHostnames() []string {
	// Read the entire file content
	data, err := os.ReadFile("data/bat_hostnames.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	// Split the content by newline into a slice of words
	words := strings.Split(string(data), "\n")

	return words
}

// Writes out similarities of one hostname to all the other
func ShowExampleSimilarities(dummyHostnames []string) {
	for i := 0; i < len(dummyHostnames); i++ {
		hostname := dummyHostnames[i]
		mh := StringToMinHash(hostname)
		dataset = append(dataset, Item{
			hostname: hostname,
			mh:       mh,
			similar:  nil,
		})
	}

	randomIndex := rand.Intn(len(dummyHostnames))
	firstItem := dataset[randomIndex].mh
	fmt.Printf("The hostname is: %s\n", dataset[randomIndex].hostname)

	sort.Sort(BySimilarity{data: dataset, mainElement: firstItem})

	fmt.Print("Similar hostanmes: ")
	for i := 0; i < len(dataset); i++ {
		s := firstItem.Similarity(dataset[i].mh)
		if s > 0 {
			fmt.Printf("%s (%f), ", dataset[i].hostname, firstItem.Similarity(dataset[i].mh))
		}
	}
	fmt.Print("\nNon-similar hostanmes: ")
	for i := 0; i < len(dataset); i++ {
		s := firstItem.Similarity(dataset[i].mh)
		if s == 0 {
			fmt.Printf("%s, ", dataset[i].hostname)
		}
	}
}

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
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

func ShowHeatmapSimilarities(hostnames []string) {
	n := len(hostnames)
	dataset = []Item{}
	for i := 0; i < len(hostnames); i++ {
		hostname := hostnames[i]
		mh := StringToMinHash(hostname)
		dataset = append(dataset, Item{
			hostname: hostname,
			mh:       mh,
			similar:  nil,
		})
	}

	similarities := make([][]float64, n)

	for i := 0; i < len(hostnames); i++ {
		similarities[i] = make([]float64, n)
		for j := 0; j < len(hostnames); j++ {
			similarities[i][j] = dataset[i].mh.Similarity(dataset[j].mh)
		}
	}
	// Save the data to external file
	file, _ := os.Create("similarity_array.csv")

	writer := csv.NewWriter(file)

	for _, row := range similarities {
		stringRow := make([]string, len(row))
		for i, value := range row {
			stringRow[i] = strconv.FormatFloat(value, 'f', -1, 64)
		}
		if err := writer.Write(stringRow); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

	writer.Flush()
	file.Close()

	cmd := exec.Command("python3", "/home/mmarjank/minhashing/plotSimilarityHeatmap.py")
	err := cmd.Run()
	fmt.Println(err)
}

func ShowHeatmapSimilaritiesNonsymmetric(hostnames []string, otherHostnames []string, title string, xAxisLabel string, yAxisLabel string) {
	n := len(hostnames)
	dataset = []Item{}
	for i := 0; i < len(hostnames); i++ {
		hostname := hostnames[i]
		mh := StringToMinHash(hostname)
		dataset = append(dataset, Item{
			hostname: hostname,
			mh:       mh,
			similar:  nil,
		})
	}

	m := len(otherHostnames)
	otherDataset := []Item{}
	for i := 0; i < len(otherHostnames); i++ {
		hostname := otherHostnames[i]
		mh := StringToMinHash(hostname)
		otherDataset = append(otherDataset, Item{
			hostname: hostname,
			mh:       mh,
			similar:  nil,
		})
	}

	similarities := make([][]float64, n)

	for i := 0; i < n; i++ {
		similarities[i] = make([]float64, m)

		for j := 0; j < m; j++ {
			similarities[i][j] = dataset[i].mh.Similarity(otherDataset[j].mh)
		}
	}

	file, _ := os.Create("similarity_plotFormatting.csv")
	writer := csv.NewWriter(file)
	if err := writer.Write([]string{title, xAxisLabel, yAxisLabel}); err != nil {
		log.Fatalln("error writing record to file", err)
	}
	writer.Flush()
	file.Close()

	// Save the data to external file
	file, _ = os.Create("similarity_array.csv")

	writer = csv.NewWriter(file)

	for _, row := range similarities {
		stringRow := make([]string, len(row))
		for i, value := range row {
			stringRow[i] = strconv.FormatFloat(value, 'f', -1, 64)
		}
		if err := writer.Write(stringRow); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

	cmd := exec.Command("python3", "/home/mmarjank/minhashing/plotSimilarityHeatmap.py")
	err := cmd.Run()
	fmt.Println(err)
}

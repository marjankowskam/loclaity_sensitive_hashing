package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func ShowSimilarityNetwork(hostnames []string) {
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

	for i := 0; i < n; i++ {
		similarities[i] = make([]float64, n)

		for j := 0; j <= i; j++ {
			similarities[i][j] = dataset[i].mh.Similarity(dataset[j].mh)
		}
	}

	// Save the data to external file
	file, _ := os.Create("data/similarity_array_network.csv")
	label_file, _ := os.Create("data/similarity_array_network_labels.csv")

	writer := csv.NewWriter(file)
	label_writer := csv.NewWriter(label_file)

	for _, row := range similarities {
		stringRow := make([]string, len(row))
		for i, value := range row {
			stringRow[i] = strconv.FormatFloat(value, 'f', -1, 64)
		}
		if err := writer.Write(stringRow); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

	if err := label_writer.Write(hostnames); err != nil {
		log.Fatal(err)
	}

	writer.Flush()
	file.Close()
	label_writer.Flush()
	label_file.Close()

	cmd := exec.Command("python3", "/home/mmarjank/minhashing/plotSimilarityNetwork.py")
	err := cmd.Run()
	fmt.Println(err)
}

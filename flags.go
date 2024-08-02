package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
)

const maxLength = 5
const flagArgRatio = 5

func getRandomFlag() string {
	flagArgPrefix := map[bool]string{false: "", true: "-"}
	isAFlag := rand.Intn(flagArgRatio) > 0
	return flagArgPrefix[isAFlag] + GetRandomString(rand.Intn(maxLength)+1)
}

func addAFlag(num int, originalFlags []string) []string {
	changedFlags := originalFlags
	for i := 0; i < num; i++ {
		j := rand.Intn(len(changedFlags) + 1) // the insertion point
		changedFlags = append(changedFlags[:j], append([]string{getRandomFlag()}, changedFlags[j:]...)...)
	}
	return changedFlags
}

func removeAFlag(num int, originalFlags []string) []string {
	if len(originalFlags) == 0 {
		return originalFlags
	}
	changedFlags := originalFlags
	for i := 0; i < num; i++ {
		j := rand.Intn(len(changedFlags)) // the deletion point
		changedFlags = append(changedFlags[:j], changedFlags[(j+1):]...)
	}
	return changedFlags
}

// GetFlags generates a sequence of num flags/arguments
func GetFlags(num int) []string {
	flagSequence := make([]string, num)

	for i := 0; i < num; i++ {
		flagSequence[i] = getRandomFlag()
	}
	return flagSequence
}

// ModifyFlags changes the flag sequence by a random addition/deletion
// num is the number of changes
func ModifyFlags(num int, originalFlags []string) []string {
	changedFlags := originalFlags
	for i := 0; i < num; i++ {
		if rand.Intn(2) < 1 {
			changedFlags = addAFlag(1, changedFlags)
		} else {
			changedFlags = removeAFlag(1, changedFlags)
		}
	}
	return changedFlags
}

// Perturb flag sequences (of lenghts from 1 to num) and note the change in similarity
func FlagExperiment(num int) ([]float64, []float64, []float64) {
	num_attempts := 100
	maxima := make([]float64, num)
	minima := make([]float64, num)
	averag := make([]float64, num)

	//j is the number or flags and arguments (our x axis)
	for j := 1; j < num; j++ {
		scores := make([]float64, num_attempts)
		for i := 0; i < num_attempts; i++ {
			a_flag := GetFlags(j)
			changed_flag := ModifyFlags(1, a_flag)
			scores[i] = StringToMinHash(strings.Join(a_flag, " ")).Similarity(StringToMinHash(strings.Join(changed_flag, " ")))
		}
		maxima[j] = slices.Max(scores)
		minima[j] = slices.Min(scores)
		averag[j] = average(scores)
	}
	PlotLines(maxima, averag, minima)
	return maxima, averag, minima
}

// because for some reason it's not a buid-in function
func average(slice []float64) float64 {
	sum := 0.0
	for _, x := range slice {
		sum += x
	}
	return sum / float64(len(slice))
}

func PlotLines(maxima []float64, average []float64, minima []float64) {
	file, _ := os.Create("data/similarity_array_lineplot.csv")

	writer := csv.NewWriter(file)

	for i := 0; i < len(maxima); i++ {
		record := []string{
			strconv.FormatFloat(maxima[i], 'f', -1, 64),
			strconv.FormatFloat(average[i], 'f', -1, 64),
			strconv.FormatFloat(minima[i], 'f', -1, 64),
		}
		if err := writer.Write(record); err != nil {
			log.Fatalf("failed writing record to file: %s", err)
		}
	}
	writer.Flush()
	file.Close()
	cmd := exec.Command("python3", "/home/mmarjank/minhashing/plotSimilarityLineplot.py")
	err := cmd.Run()
	fmt.Println(err)
}

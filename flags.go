package main

import "math/rand"

const maxLength = 5
const flagArgRatio = 5

func getRandomFlag() string {
	flagArgPrefix := map[bool]string{false: "", true: "-"}
	isAFlag := rand.Intn(flagArgRatio) > 0
	return flagArgPrefix[isAFlag] + GetRandomString(rand.Intn(maxLength)+1)
}

func GetFlags(num int) []string {
	flagSequence := make([]string, num)

	for i := 0; i < num; i++ {
		flagSequence[i] = getRandomFlag()
	}
	return flagSequence
}

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
